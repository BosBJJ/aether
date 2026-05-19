package recon

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func ScanSub(word, domain, resolver string) DNSResultAny {
	if domain == "" {
		fmt.Printf("domain string empty: %v", domain)
		return DNSResultAny{}
	}
	if word == "" {
		fmt.Printf("dictionary prefix word is empty")
		return DNSResultAny{} 
	}
	subdomain := word + "." + domain
	return QueryDNSAny(subdomain, resolver)
}

func ScanSubs(domain, wordPath, resolver string, workers int) (<- chan DNSResultAny, error) {
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
				results <- ScanSub(w, domain, resolver)
			}
		}()
	}
	
	go func() {
		defer close(word)
		defer list.Close()
		for scanner.Scan() {
			word <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("scan error: %v\n", err)
		}
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	return results, nil
}