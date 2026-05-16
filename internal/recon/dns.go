package recon

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

type DNSResult struct {
	Domain 	string 		`json:"domain"`
	A 		[]string	`json:"a"`
	AAAA 	[]string	`json:"aaaa"`
	MX 		[]string	`json:"mx"`
	NS 		[]string	`json:"ns"`
	TXT 	[]string 	`json:"txt"`
	CNAME 	[]string	`json:"cname"`
}



func QueryDNS(domain string) (DNSResult, error){
	if domain == "" {
		return DNSResult{}, fmt.Errorf("please enter a domain")
	}
	if !strings.HasSuffix(domain, ".") {
		domain = domain + "."
	}
	result := DNSResult{Domain: domain}
	
	dnsTypes := []uint16 {
		dns.TypeA,
		dns.TypeAAAA,
		dns.TypeMX,
		dns.TypeNS,
		dns.TypeTXT,
		dns.TypeCNAME,
	}
	c := new(dns.Client)
	c.Net = "tcp"
	for _, dType := range dnsTypes {
		msg := new(dns.Msg)
		msg.SetQuestion(domain, dType)
	
		res, _, err := c.Exchange(msg, "8.8.8.8:53")
		if err != nil {
			fmt.Printf("DNS query failed %d: %v", dType, err)
			continue
		}
		for _, answer := range res.Answer {
			switch rr := answer.(type) {
			case *dns.A:
				result.A = append(result.A, rr.A.String())
			case *dns.AAAA:
				result.AAAA = append(result.AAAA, rr.AAAA.String())
			case *dns.MX:
				result.MX = append(result.MX, fmt.Sprintf("%d %s", rr.Preference, rr.Mx))
			case *dns.NS:
				result.NS = append(result.NS, rr.Ns)
			case *dns.TXT:
				result.TXT = append(result.TXT, strings.Join(rr.Txt, " "))
			case *dns.CNAME:
				result.CNAME = append(result.CNAME, rr.Target)
			}
		}
	}
	return result, nil
}
