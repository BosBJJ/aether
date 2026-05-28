package http

import (
	"aether/cmd/config"
	"aether/internal/database"
	"aether/internal/httpinspect"
	"aether/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Basic HTTP page information and metadata",
	Long:  `Fetches a URL and displays basic information including status code, title, server, content type, and response time.
	JSON output truncates the body to 1000 characters.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Root().PersistentFlags().GetString("output")
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")

		fmt.Printf("looking up %s\n\n", target)

		result := httpinspect.Fetch(target, timeout)
		if result.Error != "" {
			fmt.Printf("error fetching result for %v: %v\n", target, result.Error)
			return
		}


		if config.ShouldSave() && config.DB != nil {
			data, err := json.Marshal(result)
			if err != nil {
				fmt.Printf("error marshalling scan result: %v\n", err)
				return
			}
			err = database.SaveScan(config.DB, database.ScanResult{
				CommandType: "http info",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Status: %v, Title: %v", result.Status, result.Title),
			})
			if err != nil {
				fmt.Printf("error saving scan result: %v\n", err)
				return
			}
		}



		if output == "json" {
			if len(result.Body) > 1000 {
				result.Body = result.Body[:1000]
			}
			err := utils.PrintJSON(result)
			if err != nil {
				fmt.Printf("error printing json: %v", err)
				return
			}
			return
		}


		fmt.Printf("====== HTTP/S Info for %v ======\n", target)
		fmt.Printf("Status      	: %v\n", result.Status)
		fmt.Printf("Final URL      	: %v\n", result.FinalURL)
		fmt.Printf("Title      	: %v\n", result.Title)
		fmt.Printf("Server      	: %v\n", result.Server)
		fmt.Printf("Content Type    : %v\n", result.ContentType)
		fmt.Printf("Content Size    : %v\n", len(result.Body))
		fmt.Printf("Response Time   : %v\n", result.ResponseTime)
	},
}

func init() {
	infoCmd.Flags().StringP("target", "t", "", "Domain or IP to query")
	infoCmd.Flags().BoolP("json", "j", false, "Output results as JSON")

	infoCmd.MarkFlagRequired("target")
}