package scan

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"aether/cmd/config"
	"aether/internal/database"
	"aether/internal/scanner"
	"aether/internal/utils"

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
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")
		threads, _ := cmd.Root().PersistentFlags().GetInt("threads")

		var results []scanner.PingResult
		switch  {
		case target != "":
			results = append(results, scanner.PingHost(target, time.Duration(timeout)*time.Second))
			
		case targets != "":
			hosts := strings.Split(targets, ",")
			res, err := scanner.PingHosts(hosts, time.Duration(timeout)*time.Second, threads)
			if err != nil {
				fmt.Println("error:", err)
        		return
			}
			results = append(results, res...)
		case cidr != "":
			res, err := scanner.SweepCIDR(cidr, time.Duration(timeout)*time.Second, threads)
			if err != nil {
				fmt.Println("error:", err)
        		return
			}
			results = append(results, res...)
		}

		if config.ShouldSave() && config.DB != nil {
			data, err := json.Marshal(results)
			if err != nil {
				fmt.Printf("error marshalling scan result: %v\n", err)
				return
			}
			err = database.SaveScan(config.DB, database.ScanResult{
				CommandType: "scan hosts",
				Target: target,
				RawResult: string(data),
				Summary: fmt.Sprintf("Hosts scanned: %v, Target: %v", len(results), target),
			})
			if err != nil {
				fmt.Printf("error saving scan result: %v\n", err)
				return
			}
		}



		if output == "json" {
			err := utils.PrintJSON(results)
			if err != nil {
				fmt.Printf("error marshalling json: %v\n", err)
			}
			return
		}

		for _, res := range results {
			printPingResult(res)
		}
	},
}

func init() {
	hostsCmd.Flags().StringP("target", "t", "", "Target IP or hostname")
	hostsCmd.Flags().StringP("targets", "", "", "Targets multiple IPs")
	hostsCmd.Flags().StringP("cidr", "c", "", "Scans subnet, /24 = 256.256.256.XXX, /16 = 256.256.XXX.XXX")

}

func printPingResult(result scanner.PingResult) {
	if result.Alive{
		fmt.Printf("%-20s ALIVE  latency: %.3fms\n", result.IP, result.Latency)
	} else if config.IsVerbose() {
        fmt.Printf("%-20s DEAD\n", result.IP)
    }
}