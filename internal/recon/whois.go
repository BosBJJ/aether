package recon

import (
	"fmt"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type WhoisResult struct {
	Domain      string   `json:"domain"`
	Registrar   string   `json:"registrar"`
	CreatedDate string   `json:"created_date"`
	ExpiryDate  string   `json:"expiry_date"`
	UpdatedDate string   `json:"updated_date"`
	Nameservers []string `json:"nameservers"`
	Status      []string `json:"status"`
	Raw 		string	 `json:"raw,omitempty"`
}


func Lookup(domain string) (WhoisResult, error){
	if domain == "" {
		return WhoisResult{}, fmt.Errorf("domain cannot be empty")
	}
	lookup, err := whois.Whois(domain)
	if err != nil {
		return WhoisResult{}, fmt.Errorf("whois lookup failed: %v", err)
	}
	result, err := whoisparser.Parse(lookup)
	if err != nil {
		return WhoisResult{Domain: domain, Raw: lookup}, nil
	}
	return WhoisResult{
		Domain: result.Domain.Domain,
		Registrar: result.Registrar.Name,
		CreatedDate: result.Domain.CreatedDate,
		ExpiryDate: result.Domain.ExpirationDate,
		UpdatedDate: result.Domain.UpdatedDate,
		Nameservers: result.Domain.NameServers,
		Status: result.Domain.Status,
		Raw: lookup,
	}, nil
}