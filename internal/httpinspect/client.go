package httpinspect

import (
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type HTTPResponse struct {
	URL 			string 			`json:"url"`
	FinalURL 		string			`json:"final_url,omitempty"`
	StatusCode 		int				`json:"status_code"`
	Status 			string			`json:"status"`
	Headers 		http.Header		`json:"headers"`
	Body 			string			`json:"body,omitempty"`
	Title 			string			`json:"title,omitempty"`
	Server 			string			`json:"server,omitempty"`
	ContentType 	string			`json:"content_type,omitempty"`
	ResponseTime 	time.Duration	`json:"response_time"`
	Cookies      	[]*http.Cookie	`json:"cookies,omitempty"`
	Redirects 		[]string		`json:"redirects,omitempty"`
	Error 			string			`json:"error,omitempty"`
}


func Fetch(url string, timeout int) HTTPResponse {
	if url == "" {
		return HTTPResponse{}
	}
	result := HTTPResponse{URL: url}
	var redirects []string
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			redirects = append(redirects, req.URL.String())
			return nil
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		result.Error = err.Error()
		return result
	}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		result.Error = err.Error()
		return result
	}
	defer resp.Body.Close()
	result.ResponseTime = time.Since(start)
	result.Redirects = redirects
	result.FinalURL = resp.Request.URL.String()

	result.StatusCode = resp.StatusCode
	result.Status = resp.Status
	result.Headers = resp.Header
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Error = err.Error()
		return result
	}
	result.Body = string(body)
	result.Title = extractTitle(result.Body)
	result.Server = resp.Header.Get("Server")
	result.ContentType = resp.Header.Get("Content-Type")
	result.Cookies = resp.Cookies()


	return result

}


func extractTitle(body string) string {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return ""
	}
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.Title {
			if node.FirstChild != nil {
				return node.FirstChild.Data
			}
		}
	}
	return ""
}