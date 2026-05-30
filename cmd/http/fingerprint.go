package http

import (
	"github.com/BosBJJ/aether/cmd/config"
	"github.com/BosBJJ/aether/internal/database"
	"github.com/BosBJJ/aether/internal/httpinspect"
	"github.com/BosBJJ/aether/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var fingerprintCmd = &cobra.Command{
	Use:   "fingerprint",
	Aliases: []string{"fprint", "fp"},
	Short: "Fingerprint the technology stack of a web server",
	Long:  `Fetches a URL and attempts to identify the underlying technology stack
by analyzing HTTP response headers, cookies, and page content.

Outputs detected technologies grouped by category including web server, CDN,
backend, frontend, CMS, analytics, and error tracking.

Use --json to output results as JSON.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Root().PersistentFlags().GetString("output")
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")


		result := httpinspect.GetTechStack(target, timeout)
		if result == nil {
			fmt.Printf("error fetching headers for %v\n", target)
			return
		}

		var out []TechStackJSON
		for _, sig := range result {
			out = append(out, TechStackJSON{Name: sig.Name, Category: sig.Category})
		}

		if config.ShouldSave() && config.DB != nil {
			data, err := json.Marshal(out)
			if err != nil {
				fmt.Printf("error marshalling results: %v",err)
			}
			err = database.SaveScan(config.DB, database.ScanResult{
				CommandType: "http fingerprint",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Detected %v technologies", len(out)),
			})
			if err != nil {
				fmt.Printf("error saving scan result: %v\n", err)
				return
			}
		}

		if output == "json" {
			err := utils.PrintJSON(out)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
				return
			}
			return
		}
		printOrder := []string{"Web Server", "CDN", "Backend", "Frontend", "CMS", "Analytics", "Error Tracking"}
		
		fmt.Println()
		fmt.Printf("TECH STACK FOR %v\n", target)
		fmt.Println("========================")
		for _, order := range printOrder {
			for _, sig := range result {
				if sig.Category == order {
					fmt.Printf("%-20v : %v\n", sig.Category, sig.Name)
				}
			}
		}

	},
}

func init() {
	fingerprintCmd.Flags().StringP("target", "t", "", "Domain or IP to query")

	fingerprintCmd.MarkFlagRequired("target")
}


type TechStackJSON struct {
	Category 	string   	`json:"category"`
	Name    	string 		`json:"name"`
}