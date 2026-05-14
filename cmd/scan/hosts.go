package scan

import (
	"fmt"
	"strings"
	"time"

	"aether/cmd/config"
	"aether/internal/scanner"

	"github.com/spf13/cobra"
)

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Discover live hosts on a network",
	Long:  `Performs host discovery (ping sweep) on a subnet or list of IPs.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		targets, _ := cmd.Flags().GetString("targets")
		cidr, _ := cmd.Flags().GetString("cidr")

		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")
		threads, _ := cmd.Root().PersistentFlags().GetInt("threads")

		switch  {
		case target != "":
			result := scanner.PingHost(target, time.Duration(timeout)*time.Second)
			printPingResult(result)
		case targets != "":
			hosts := strings.Split(targets, ",")
			results, err := scanner.PingHosts(hosts, time.Duration(timeout)*time.Second, threads)
			if err != nil {
				fmt.Println("error:", err)
        		return
			}
			for _, result := range results {
				printPingResult(result)
			}
		case cidr != "":
			results, err := scanner.SweepCIDR(cidr, time.Duration(timeout)*time.Second, threads)
			if err != nil {
				fmt.Println("error:", err)
        		return
			}
			for _, result := range results {
				printPingResult(result)
			}
		}
		fmt.Println("Scan complete")
	},
}

func init() {
	hostsCmd.Flags().StringP("target", "t", "", "Target IP or hostname")
	hostsCmd.Flags().StringP("targets", "", "", "Targets multiple IPs")
	hostsCmd.Flags().StringP("cidr", "c", "", "Scans subnet, /24 = 256.256.256.XXX, /16 = 256.256.XXX.XXX")

}

func printPingResult(result scanner.PingResult) {
	if result.Alive{
		fmt.Printf("%-20s ALIVE  latency: %.3fms\n", result.IP, result.Latency.Seconds()*1000)
	} else if config.IsVerbose() {
        fmt.Printf("%-20s DEAD\n", result.IP)
    }
}