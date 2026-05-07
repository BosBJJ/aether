/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
    outputFormat string
    saveToDB     bool
    threads      int
    timeout      int
    quiet        bool
    verbose      bool
)

// rootCmd represents the base command when called without any subcommands
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aether.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, json, html)")
	rootCmd.PersistentFlags().BoolVar(&saveToDB, "save", false, "Save scan results to database")
	rootCmd.PersistentFlags().IntVar(&threads, "threads", 50, "Number of concurrent threads to use")
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 5, "Timeout in seconds for network operations")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode - show minimal output")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode - show detailed output")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func GetOutputFormat() string {
    return outputFormat
}

func ShouldSave() bool {
    return saveToDB
}

func GetThreads() int {
    if threads <= 0 {
        return 50
    }
    return threads
}

func GetTimeout() int {
    if timeout <= 0 {
        return 5
    }
    return timeout
}

func IsQuiet() bool {
    return quiet
}

func IsVerbose() bool {
    return verbose
}