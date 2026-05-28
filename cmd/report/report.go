package report

import "github.com/spf13/cobra"

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "View and manage scan reports",
	Long:  `Report provides tools to list, view, and generate reports from saved scan results.`,
}

func Init() {
	reportCmd.AddCommand(
		listCmd,
		viewCmd,
		deleteCmd,
		exportCmd,
	)
}

func GetReportCmd() *cobra.Command {
	Init()
	return reportCmd
}
