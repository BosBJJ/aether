package recon

import (
	"aether/internal/recon"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var passiveCmd = &cobra.Command{
	Use:   "passive",
	Short: "Passive recon using certificate transparency logs",
	Long:  `Queries certificate transparency logs to discover subdomains associated with a target domain. No active probing is performed.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		jsonOut, _ := cmd.Flags().GetBool("json")
		limit, _ := cmd.Flags().GetInt("limit")
		fmt.Printf("looking up %s\n\n", target)

		result, err := recon.PassiveRecon(target, limit)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		
		if jsonOut {
			data, err := json.Marshal(result)
			if err != nil {
				fmt.Printf("error marshalling json: %v\n", err)
        		return
			}
			fmt.Println(string(data))
    		return
		}

		fmt.Printf("====== PASSIVE RECON RESULT FOR %v ======\n", target)
		fmt.Printf("Source used: %v\n", result.Source)
		if result.TotalCerts == 0 {
			fmt.Println("no certificates found, please try again")
			return
		}
		fmt.Printf("Certificates found: %v\n", result.TotalCerts)
		fmt.Println("Unique subdomains: ")
		for _, sub := range result.Subdomains {
			fmt.Printf("-     %v\n", sub)
		}
		fmt.Printf("Oldest cert: %v\n", result.FirstSeen)
		fmt.Printf("Most recent cert: %v\n", result.LastSeen)
	},
}

func init() {
	passiveCmd.Flags().StringP("target", "t", "", "Domain or IP to query")
	passiveCmd.Flags().BoolP("json", "j", false, "Output results as JSON")
	passiveCmd.Flags().IntP("limit", "l", 100, "Maximum number of certificate records to fetch")

	passiveCmd.MarkFlagRequired("target")
}