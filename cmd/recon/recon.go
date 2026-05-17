package recon

import "github.com/spf13/cobra"

var reconCmd = &cobra.Command{
	Use:   "recon",
	Short: "Domain, DNS, and passive reconnaissance tools",
	Long:  `Recon includes WHOIS lookups, DNS enumeration, subdomain discovery,
and passive intelligence gathering for domains and internet-facing infrastructure.`,
}

func Init() {
	reconCmd.AddCommand(
		whoisCmd,
		dnsCmd,
		subCmd,
	)
}

func GetReconCmd() *cobra.Command {
	Init()
	return reconCmd
}