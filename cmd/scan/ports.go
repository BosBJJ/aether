package scan

import (
	"fmt"
	"time"

	"aether/internal/scanner"
	"github.com/spf13/cobra"
)

var portsCmd = &cobra.Command{
	Use:   "ports",
	Short: "Scan ports on a target",
	Long:  `Performs a TCP port scan on a target. Defaults to common ports (80, 443, 22, 21, 25) unless --top100, --top1000, or --all is specified `,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		top100, _ := cmd.Flags().GetBool("top100")
		top1000, _ := cmd.Flags().GetBool("top1000")
		all, _ := cmd.Flags().GetBool("all")
		timeout, _ := cmd.Flags().GetInt("timeout")
		threads, _ := cmd.Flags().GetInt("threads")
		verbose, _ := cmd.Flags().GetBool("verbose")


		var ports []int
		switch {
		case all:
			ports = make([]int, 65535)
			for i := range ports {
				ports[i] = i + 1
			}
		case top1000:
			ports = scanner.Top1000Ports
		case top100:
			ports = scanner.Top100Ports
		default:
			ports = []int{80, 443, 22, 21, 25}
		}
		
		results, err := scanner.ScanPorts(target, ports, time.Duration(timeout)*time.Second, threads)
		if err != nil {
			fmt.Printf("error: %v\n", err)
   			 return
		}
		fmt.Printf("Scan completed for %s\n", target)
		for _, result := range results {
			if result.State == "Open" || verbose {
				fmt.Printf("%-6d %s\n", result.Port, result.State)
			}
		}
	},
}

func init() {
	portsCmd.Flags().StringP("target", "t", "", "Target IP or hostname (required)")
	portsCmd.Flags().Bool("all", false, "Port range 1-65535")
	portsCmd.Flags().Bool("top100", false, "Use top 100 common ports")
	portsCmd.Flags().Bool("top1000", false, "Use top 1000 common ports")

	portsCmd.MarkFlagRequired("target")
}