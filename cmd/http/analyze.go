package http

import (
	"aether/internal/httpinspect"
	"aether/internal/utils"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a URL for suspicious parameters and open redirect risks",
	Long:  `Fetches a URL and inspects it for security concerns including:
- Suspicious or sensitive query parameters (e.g. redirect, token, next)
- Suspicious TLDs commonly associated with phishing
- Open redirect indicators across the full redirect chain

Use --json to output results in JSON format.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		jsonOut, _ := cmd.Flags().GetBool("json")
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")

		result, err := httpinspect.AnalyzeURL(target, timeout)
		if err != nil {
			fmt.Printf("error analyzing headers for %v: %v\n", target, err)
			return
		}


		if jsonOut {
			err := utils.PrintJSON(result)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
				return
			}
			return
		}
		
		
		fmt.Println()
		fmt.Printf("Analysis Results for %v:\n", result.OriginalURL)
		fmt.Println("=============================")
		fmt.Printf("Final URL: %v\n", result.FinalURL)
		oHost, _ := url.Parse(result.OriginalURL)
		fHost, _ := url.Parse(result.FinalURL)
		if oHost.Hostname() != fHost.Hostname() {
			fmt.Printf("⚠ Domain changed: %v -> %v\n", oHost.Hostname(), fHost.Hostname())
		}
		fmt.Printf("Number of redirects : %v\n", len(result.Redirects))
		if len(result.Findings) > 0 {
			fmt.Println()
			fmt.Println("Findings")
			for _, f := range result.Findings {
				risk := fmt.Sprintf("[%s RISK]", strings.ToUpper(f.Severity))
				fmt.Printf("%-16v: %v\n", risk, f.Description)
			}
		}
		fmt.Println()
		fmt.Printf("Risk Level: %v\n", result.RiskLevel)
	},
}

func init() {
	analyzeCmd.Flags().StringP("target", "t", "", "Domain or IP to query")
	analyzeCmd.Flags().BoolP("json", "j", false, "Output results as JSON")

	analyzeCmd.MarkFlagRequired("target")
}