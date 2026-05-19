package recon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PassiveResult struct {
    Domain        string 		`json:"domain"`
	Source		  string		`json:"source"`
    Subdomains    []string 		`json:"subdomains"`
    TotalCerts    int			`json:"certificates_found"`
    FirstSeen     string 		`json:"first_seen,omitempty"`
    LastSeen      string 		`json:"last_seen,omitempty"`
}

func PassiveRecon(domain string) (PassiveResult, error) {
	if domain == "" {
		return PassiveResult{}, fmt.Errorf("domain empty")
	}

	url := "https://crt.sh/?q=%25." + domain + "&output=json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PassiveResult{}, fmt.Errorf("failed to make request: %v", err)
	}
	req.Header.Set("User-Agent", "aether-recon/1.0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PassiveResult{}, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PassiveResult{}, fmt.Errorf("crt.sh returned status %d", resp.StatusCode)
	}

	type crtResp struct {
		CommonName string `json:"common_name"`
		NameValue string `json:"name_value"`
		NotBefore string `json:"not_before"`
		NotAfter string `json:"not_after"`
	}
	var results []crtResp

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&results)
	if err != nil {
		return PassiveResult{}, fmt.Errorf("parsing error: %v", err)
	}
	
	var subdomains []string
	tracked := map[string]bool{}
	first, last := "", ""
	for _, result := range results {
		names := strings.Split(result.NameValue, "\n")
		for _, name := range names {
			if strings.HasPrefix(name, "*.") {
				continue
			}
			if strings.Contains(name, "@") {
				continue
			}
			if !strings.HasSuffix(name, domain) {
				continue
			}
			if !tracked[name] {
				subdomains = append(subdomains, name)
				tracked[name] = true
			}
		}
		if first == "" {
			first = result.NotBefore
		}
		if last == "" {
			last = result.NotAfter
		}
		if result.NotBefore < first {
			first = result.NotBefore
		}
		if result.NotAfter > last {
			last = result.NotAfter
		}
	}


	return PassiveResult{
		Domain: domain,
		Source: "crt.sh",
		Subdomains: subdomains,
		TotalCerts: len(results),
		FirstSeen: first,
		LastSeen: last,
	}, nil
}