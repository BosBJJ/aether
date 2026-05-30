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

var headerCmd = &cobra.Command{
	Use:   "headers",
	Aliases: []string{"header"},
	Short: "Analyze HTTP security and information-leakage headers",
	Long:  `Fetches a URL and evaluates its HTTP response headers against a set of
known security and information-leakage checks.

Outputs a categorized report showing which headers are properly configured,
missing, or leaking sensitive information, along with a summary and recommendations.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Root().PersistentFlags().GetString("output")
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")


		result := httpinspect.AnalyzeHeaders(target, timeout)
		if result == nil {
			fmt.Printf("error analyzing headers for %v\n", target)
			return
		}

		var out []HeaderResultJSON
		for _, h := range result {
			out = append(out, HeaderResultJSON{
				Name:    h.Name,
				Present: h.Present,
				Value:   h.Value,
			})
		}

		var passed, missing, leaking int
		var secHeaders []httpinspect.HeaderResult
		var leakHeaders []httpinspect.HeaderResult

		for _, h := range result {
			if httpinspect.IsLeakHeader(h.Name){
				leakHeaders = append(leakHeaders, h)
			} else {
				secHeaders = append(secHeaders, h)
			}
		}

		for _, h := range secHeaders {
			if h.Present {
				passed++
			} else {
				missing++
			}
		}

		for _, h := range leakHeaders {
			if h.Present {
				leaking++
			}
		}

		if config.ShouldSave() && config.DB != nil {
			data, err := json.Marshal(out)
			if err != nil {
				fmt.Printf("error marshalling results: %v",err)
			}
			err = database.SaveScan(config.DB, database.ScanResult{
				CommandType: "http headers",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Headers: %v secure, %v missing, %v leaking", passed, missing, leaking),
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
		
		
		fmt.Println()
		fmt.Println("Security Headers:")
		fmt.Println("==================")
		for _, h := range secHeaders {
			symbol := "⚠"
			value := "Missing"
			if h.Present {
				symbol = "✓"
    			value = h.GoodLabel
			} 
			fmt.Printf("%v %-40v : %v\n", symbol, h.DisplayName, value)
		}

		if len(leakHeaders) > 0 {
			fmt.Println()
			fmt.Println("Info Leakage:")
			fmt.Println("==================")
			for _, h := range leakHeaders {
				symbol := "✓"
				value := "Hidden"
				if h.Present {
					symbol = "⚠"
					value = h.Value
				}
				fmt.Printf("%v %-40v : %v\n", symbol, h.DisplayName, value)
			}
		}

		fmt.Println()
		fmt.Println("Summary:")
		fmt.Printf("-- %v properly configured headers\n", passed)
		fmt.Printf("-- %v missing or weak headers\n", missing)
		fmt.Printf("-- %v headers leaking information\n", leaking)

		if missing == 0 && leaking == 0 {
			return
		}
		fmt.Println()
		fmt.Println("Recommendations:")
		for _, h := range secHeaders {
			if !h.Present {
				fmt.Printf("-- %v\n", h.Recommendation)
			}
		}
		for _, h := range leakHeaders {
			if h.Present {
				fmt.Printf("-- %v\n", h.Recommendation)
			}
		}

	},
}

func init() {
	headerCmd.Flags().StringP("target", "t", "", "Domain or IP to query")

	headerCmd.MarkFlagRequired("target")
}

type HeaderResultJSON struct {
	Name    string `json:"name"`
	Present bool   `json:"present"`
	Value   string `json:"value,omitempty"`
}