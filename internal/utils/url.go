package utils

import "strings"


func NormalizeURL(url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") { 
		return url 
	}
	return "https://" + url
}