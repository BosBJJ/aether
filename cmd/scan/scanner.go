package scan

import "github.com/spf13/cobra"

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Host and Port scanning utilities",
	Long:  `Scan provides tools for host discovery, port scanning, banner grabbing, and basic service identification.`,
}

func Init() {
	scanCmd.AddCommand(
		portsCmd,
		hostsCmd,
		bannersCmd,
		serviceCmd,
	)
}

func GetScanCmd() *cobra.Command {
	Init()
	return scanCmd
}