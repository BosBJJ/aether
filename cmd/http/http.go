package http

import "github.com/spf13/cobra"

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Analyze and inspect HTTP targets",
	Long:  `Perform HTTP-based reconnaissance including header analysis, 
technology fingerprinting, security configuration checks, and deep URL inspection.`,
}

func Init() {
	httpCmd.AddCommand(
		infoCmd,
		headerCmd,
	)
}

func GetHTTPCmd() *cobra.Command {
	Init()
	return httpCmd
}