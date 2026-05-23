package http

import (
	"aether/internal/httpinspect"
	"aether/internal/utils"
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
		jsonOut, _ := cmd.Flags().GetBool("json")
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")

		if !jsonOut {
			fmt.Printf("looking up %s\n\n", target)
		}

		result := httpinspect.GetTechStack(target, timeout)
		if result == nil {
			fmt.Printf("error fetching headers for %v\n", target)
			return
		}

		var out []TechStackJSON
		for _, sig := range result {
			out = append(out, TechStackJSON{Name: sig.Name, Category: sig.Category})
		}

		if jsonOut {
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
	fingerprintCmd.Flags().BoolP("json", "j", false, "Output results as JSON")

	fingerprintCmd.MarkFlagRequired("target")
}


type TechStackJSON struct {
	Category 	string   	`json:"category"`
	Name    	string 		`json:"name"`
}