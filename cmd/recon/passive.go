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

var passiveCmd = &cobra.Command{
	Use:   "passive",
	Short: "Passive recon using certificate transparency logs",
	Long:  `Queries certificate transparency logs to discover subdomains associated with a target domain. No active probing is performed.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Root().PersistentFlags().GetString("output")
		limit, _ := cmd.Flags().GetInt("limit")

		result, err := recon.PassiveRecon(target, limit)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		if config.ShouldSave() && config.DB != nil {
			data, err := json.Marshal(result)
			if err != nil {
				fmt.Printf("error marshalling scan result: %v\n", err)
				return
			}
			err = database.SaveScan(config.DB, database.ScanResult{
				CommandType: "recon passive",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Domain :%v, Certificates: %v", result.Domain, result.TotalCerts),
			})
			if err != nil {
				fmt.Printf("error saving scan result: %v\n", err)
				return
			}
		}

		
		if output == "json" {
			err := utils.PrintJSON(result)
			if err != nil {
				fmt.Printf("error marshalling json: %v\n", err)
        		return
			}
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
	passiveCmd.Flags().IntP("limit", "l", 100, "Maximum number of certificate records to fetch")

	passiveCmd.MarkFlagRequired("target")
}