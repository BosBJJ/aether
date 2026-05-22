package httpinspect

import "fmt"


type HeaderResult struct {
    HeaderCheck
    Present bool   `json:"present"`
    Value   string `json:"value,omitempty"`
}


func AnalyzeHeaders(url string, timeout int) []HeaderResult {
	data := Fetch(url, timeout)
	if data.Error != "" {
		fmt.Printf("error: %v", data.Error)
	}
	var results []HeaderResult
	for _, check := range SecurityHeaders {
		headerRes := HeaderResult{HeaderCheck: check}
		value := data.Headers.Get(check.Name)
		if value != "" {
			headerRes.Present = true
			headerRes.Value = value
		}
		results = append(results, headerRes)
	}
	return results
}

func IsLeakHeader(name string) bool {
	switch name {
	case "Server", "X-Powered-By", "X-AspNet-Version":
		return true
	}
	return false
}