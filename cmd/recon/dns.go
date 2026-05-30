package recon

import (
	"github.com/BosBJJ/aether/cmd/config"
	"github.com/BosBJJ/aether/internal/database"
	"github.com/BosBJJ/aether/internal/recon"
	"github.com/BosBJJ/aether/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Perform DNS reconnaissance on a domain",
	Long:  `Queries A, AAAA, MX, NS, TXT, and CNAME records for a given domain using Google's public DNS resolver.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		resolver, _ := cmd.Flags().GetString("resolver")
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		if output == "table" {
			fmt.Printf("looking up %v\n\n", target)
		}

		result, err := recon.QueryDNS(target, resolver)
		if err != nil {
			fmt.Printf("error :%v\n", err)
			return
		}

		if config.ShouldSave() && config.DB != nil {
			data, err := json.Marshal(result)
			if err != nil {
				fmt.Printf("error marshalling scan result: %v\n", err)
				return
			}
			err = database.SaveScan(config.DB, database.ScanResult{
				CommandType: "recon dns",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Domain: %v, Records: %v", result.Domain, len(result.A)+len(result.AAAA)+len(result.MX)+len(result.NS)+len(result.TXT)+len(result.CNAME)),
			})
			if err != nil {
				fmt.Printf("error saving scan result: %v\n", err)
				return
			}
		}



		if output == "json" {
			err := utils.PrintJSON(result)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
				return
			}
			return
		}

		fmt.Println("========DNS RESULTS========")
		fmt.Printf("Domain:      %v\n\n", result.Domain)
		fmt.Println("A:")
		for _, a := range result.A {
			fmt.Printf("    -%v\n", a)
		}
		fmt.Println("\nAAAA:")
		for _, r := range result.AAAA {
			fmt.Printf("    -%v\n", r)
		}
		fmt.Println("\nMX:")
		for _, r := range result.MX {
			fmt.Printf("    - %v\n", r)
		}
		fmt.Println("\nNS:")
		for _, r := range result.NS {
			fmt.Printf("    - %v\n", r)
		}
		fmt.Println("\nTXT:")
		for _, r := range result.TXT {
			fmt.Printf("    - %v\n", r)
		}
		fmt.Println("\nCNAME:")
		for _, r := range result.CNAME {
			fmt.Printf("    - %v\n", r)
		}
			
	},
}

func init() {
	dnsCmd.Flags().StringP("target", "t", "", "Domain or IP to query")
	dnsCmd.Flags().StringP("resolver", "r", "8.8.8.8:53", "DNS resolver to use")
	dnsCmd.MarkFlagRequired("target")
}