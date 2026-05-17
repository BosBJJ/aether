package recon

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func ScanSub(word, domain string) DNSResultAny {
	if domain == "" {
		fmt.Printf("domain string empty: %v", domain)
		return DNSResultAny{}
	}
	if word == "" {
		fmt.Printf("dictionary prefix word is empty")
		return DNSResultAny{} 
	}
	subdomain := word + "." + domain
	return QueryDNSAny(subdomain)
}

func ScanSubs(domain, wordPath string, workers int) (<- chan DNSResultAny, error) {
	list, err := os.Open(wordPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open word list: %v", err)
	}
	
	var wg sync.WaitGroup
	results := make(chan DNSResultAny, workers)
	word := make(chan string)
	scanner := bufio.NewScanner(list)
	
	for range workers {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			for w := range word {
				results <- ScanSub(w, domain)
			}
		}()
	}
	
	go func() {
		for scanner.Scan() {
			word <- scanner.Text()
		}
		close(word)
		list.Close()
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	return results, nil
}