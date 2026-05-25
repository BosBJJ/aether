package recon

import (
	"aether/internal/recon"
	"aether/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "substring",
	Aliases: []string{"subs", "sub"},
	Short: "Enumerate subdomains using a DNS wordlist",
	Long:  `Performs active subdomain enumeration by querying DNS for common subdomain prefixes.
Supports small (5k), medium (110k), and large (1M) wordlists with configurable concurrency.
Automatically detects wildcard DNS to avoid false positives.
Use --threads to control concurrency (default 50).`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		threads, _ := cmd.Root().PersistentFlags().GetInt("threads")
		size, _ := cmd.Flags().GetString("size")
		resolver, _ := cmd.Flags().GetString("resolver")
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		path := ""
		numWords := ""
		switch size {
		case "small":
			path = "data/wordlists/subdomains-top1million-5000.txt"
			numWords = "5000"
		case "medium":
			path = "data/wordlists/subdomains-top1million-110000.txt"
			numWords = "110000"
		case "large":
			path = "data/wordlists/subdomains-top1million-full.txt"
			numWords = "1 million"
		default:
			path = "data/wordlists/subdomains-top1million-5000.txt"
			numWords = "5000"
		}

		wildcardTest := recon.ScanSub("thisshouldneverwork", target, resolver)
		if wildcardTest.Response {
			fmt.Println("Domain using wildcard")
			return
		}

		if output == "table" {
			fmt.Println("Domain not using wildcard")
			fmt.Printf("Preparing scan for domain: %v with %v words\n",target, numWords)
		}

		res, err := recon.ScanSubs(target, path, resolver, threads)
		if err != nil {
			fmt.Printf("error: %v\n", err)
    		return
		}

		var hits []recon.DNSResultAny
		for r := range res {
			if r.Response {
				if output == "table" {
					fmt.Printf("--   %v\n",r.Domain)
				}
				hits = append(hits, r)
			}
		}

		if output == "json" {
			err := utils.PrintJSON(hits)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
			}
			return
		}

		
		if len(hits) == 0 {
			fmt.Printf("No subdomains found for: %v\n", target)
		}
		
			
	},
}

func init() {
	subCmd.Flags().StringP("target", "t", "", "Domain or IP to query")
	subCmd.Flags().StringP("size", "s", "small", "Wordlist size: small, medium, large")
	subCmd.Flags().StringP("resolver", "r", "8.8.8.8:53", "DNS resolver to use")
	subCmd.MarkFlagRequired("target")

}