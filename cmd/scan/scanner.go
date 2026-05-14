package scan

import "github.com/spf13/cobra"

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Host and Port scanning utilities",
	Long:  `Host discovery, port scanning, banner grabbing, and service detection.`,
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