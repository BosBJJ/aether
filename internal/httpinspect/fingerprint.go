package httpinspect

import (
	"fmt"
	"strings"
)


func GetTechStack(url string, timeout int) []TechSignature {
	var techstack []TechSignature
	
	data := Fetch(url, timeout)
	if data.Error != "" {
		fmt.Println("error: ", data.Error)
		return nil
	}
	for _, sig := range FingerprintDB {
		var found bool
		for headerName, value := range sig.Headers {
			if strings.Contains(strings.ToLower(data.Headers.Get(headerName)), strings.ToLower(value)) {
				found = true
			}
		}
		for _, value := range sig.BodyPattern {
			if strings.Contains(strings.ToLower(data.Body), strings.ToLower(value)) {
				found = true
			}
		}
		for _, sigVal := range sig.Cookies {
			for _, cookie := range data.Cookies {
				if strings.Contains(strings.ToLower(cookie.Name), strings.ToLower(sigVal)) {
					found = true
				}
			}
		}
		if found {
			techstack = append(techstack, sig)
		}
	}

	mostLikely := map[string]TechSignature{}
	for _, sig := range techstack {
		existing, ok := mostLikely[sig.Category]
		if !ok || existing.Confidence < sig.Confidence {
			mostLikely[sig.Category] = sig
		}
	}
	var result []TechSignature
	for _, sig := range mostLikely {
		result = append(result, sig)
	}
	
	return result
}