package recon

import (
	"aether/cmd/config"
	"aether/internal/database"
	"aether/internal/recon"
	"aether/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "Look up WHOIS information for a domain or IP",
	Long:  `Performs a WHOIS lookup to retrieve registration and ownership information such as registrar, creation date, expiration date, and nameservers.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		if output == "table" {
			fmt.Printf("looking up %s\n\n", target)
		}
		
		result, err := recon.Lookup(target)
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
				CommandType: "recon whois",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Domain: %v, Registrar: %v, Created: %v, Expires: %v", result.Domain, result.Registrar, result.CreatedDate, result.ExpiryDate),
			})
			if err != nil {
				fmt.Printf("error saving scan result: %v\n", err)
				return
			}
		}



		
		if output == "json" {
			err = utils.PrintJSON(result)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
			}
			return
		}
		
		fmt.Println("=====WHOIS RESULTS=====")
		fmt.Printf("Domain:      %v\n", result.Domain)
		fmt.Printf("Registrar:   %v\n", result.Registrar)
		fmt.Printf("Created:     %v\n", result.CreatedDate)
		fmt.Printf("Updated:     %v\n", result.UpdatedDate)
		fmt.Printf("Expires:     %v\n", result.ExpiryDate)
		fmt.Println("Nameservers:")
		for _, serv := range result.Nameservers {
			fmt.Printf("    -%v\n", serv)
		}
		fmt.Println("Statuses:")
		for _, stat := range result.Status {
			fmt.Printf("    -%v\n", stat)
		}
			
	},
}

func init() {
	whoisCmd.Flags().StringP("target", "t", "", "Domain or IP to query")
	whoisCmd.MarkFlagRequired("target")
}