package httpinspect

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)


func AnalyzeURL(url string, timeout int) (URLAnalysis, error) {
	data := Fetch(url, timeout)
	if data.Error != "" {
		return URLAnalysis{}, fmt.Errorf("error fetching data for %v\n", url)
	}
	
	var findings []Finding

	allURLs := append(data.Redirects, data.URL, data.FinalURL)
	for _, u := range allURLs {
		if res, err := scanURL(u); err == nil {
			findings = append(findings, res...)
		}
	}

	res := scanBody(data.Body, data.FinalURL)
	if res != nil {
		findings = append(findings, res...)
	}

	riskLevel := "None"
	for _, f := range findings {
		if f.Severity == "High" {
			riskLevel = "High"
			break
		}
		if f.Severity == "Medium" {
			riskLevel = "Medium"
		}
	}

	
	return URLAnalysis{
		OriginalURL: data.URL,
		FinalURL: data.FinalURL,
		Redirects: data.Redirects,
		Findings: findings,
		RiskLevel: riskLevel,
	}, nil
	
}

func scanURL(toScan string) ([]Finding, error) {
	var results []Finding
	u, err := url.Parse(toScan)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	for param, data := range u.Query() {
		found := Finding{URL: toScan}
		if SuspiciousParams[param] != "" {
			found.Severity = SuspiciousParams[param]
			if len(data) > 1 {
				found.Description = fmt.Sprintf("suspicious param %q found on %s with values: %s", param, found.URL, strings.Join(data, ", "))
			} else {
				found.Description = fmt.Sprintf("suspicious param %q found on %s with value: %s", param, found.URL, data[0])
			}
			results = append(results, found)
		}
	}
	hostname := strings.ToLower(u.Hostname())
	for _, tld := range SuspiciousTLDs {
		found := Finding{URL: toScan}
		if strings.HasSuffix(hostname, tld) {
			found.Description = fmt.Sprintf("suspicious TLD found: %q", hostname)
			found.Severity = "High"
			results = append(results, found)
		}
	}
	return results, nil
}

func scanBody(body string, baseURL string) []Finding {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil
	}
	var results []Finding
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					hostname, _ := url.Parse(attr.Val)
					for _, tld := range SuspiciousTLDs {
						if strings.HasSuffix(hostname.Hostname(), tld){
							results = append(results, Finding{
								URL: baseURL,
								Description: fmt.Sprintf("suspicious TLD found: %q", hostname.Hostname()),
								Severity: "High",
							})
						}
					}
				}
			}
		}
		if node.Type == html.ElementNode && node.DataAtom == atom.Iframe {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					results = append(results, Finding{
						URL: baseURL,
						Description: fmt.Sprintf("suspicious iframe found with src: %q", attr.Val),
						Severity: "High",
					})
				}
			}
		}
		if node.Type == html.ElementNode && node.DataAtom == atom.Form {
			for _, attr := range node.Attr {
				if attr.Key == "action" {
					base, _ := url.Parse(baseURL)
					dest, _ := url.Parse(attr.Val)
					resolved :=base.ResolveReference(dest)
					if resolved.Hostname() != base.Hostname() {
						results = append(results, Finding{
							URL: baseURL,
							Description: fmt.Sprintf("form submits to external domain: %q", resolved.Hostname()),
							Severity: "High",
						})
					}
				}
			}
		}
	}
	return results
}