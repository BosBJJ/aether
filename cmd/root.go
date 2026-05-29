/*
Copyright © 2026 https://github.com/BosBJJ/aether
*/
package cmd

import (
	"aether/cmd/config"
	"aether/cmd/crypto"
	"aether/cmd/http"
	"aether/cmd/recon"
	"aether/cmd/report"
	"aether/cmd/scan"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aether",
	Short: "A Go-based CLI toolkit for security auditing and network reconnaissance",
	Long: `Aether is a clean, modular, and production-oriented command-line application designed to assist with security auditing, network discovery, and infrastructure assessment. 
	Built entirely in Go, it demonstrates strong backend development practices through high-performance tools, secure coding, and clean architecture.
	
	This project is intended for educational purposes and authorized security testing only.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Aether! Use --help to see available commands.")
	},
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// ====================== GLOBAL FLAGS ======================
	rootCmd.PersistentFlags().StringVarP(&config.OutputFormat, "output", "o", "table", "Output format (table, json, html)")
	rootCmd.PersistentFlags().BoolVar(&config.SaveToDB, "save", false, "Save scan results to database")
	rootCmd.PersistentFlags().IntVar(&config.Threads, "threads", 50, "Number of concurrent threads to use")
	rootCmd.PersistentFlags().IntVar(&config.Timeout, "timeout", 5, "Timeout in seconds for network operations")
	rootCmd.PersistentFlags().BoolVarP(&config.Quiet, "quiet", "q", false, "Quiet mode - show minimal output")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "Verbose mode - show detailed output")
	rootCmd.PersistentFlags().BoolVarP(&config.Pretty, "pretty", "p", false, "Pretty print JSON output")
	
	// ====================== REGISTER COMMANDS ======================

	rootCmd.AddCommand(crypto.GetCryptoCmd())
	rootCmd.AddCommand(scan.GetScanCmd())
	rootCmd.AddCommand(recon.GetReconCmd())
	rootCmd.AddCommand(http.GetHTTPCmd())
	rootCmd.AddCommand(report.GetReportCmd())
}
