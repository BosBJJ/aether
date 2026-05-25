package scan

import (
	"fmt"
	"time"

	"aether/internal/scanner"
	"aether/internal/utils"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Full service scan (ports + banners)",
	Long:  `Connects to a specified port on a target host and grabs the service banner. Supports HTTP, HTTPS, SSH, FTP, SMTP, IMAP, and POP3.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		top100, _ := cmd.Flags().GetBool("top100")
		top1000, _ := cmd.Flags().GetBool("top1000")
		all, _ := cmd.Flags().GetBool("all")
		
		var ports []int
		var openPorts []int
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")
		threads, _ := cmd.Root().PersistentFlags().GetInt("threads")
		output, _ := cmd.Root().PersistentFlags().GetString("output")

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
		if output == "table" {
			fmt.Printf("Starting service scan on %s (%d ports)...\n", target, len(ports))
		}
		scannedPorts, err := scanner.ScanPorts(target, ports, time.Duration(timeout)*time.Second, threads)
		if err != nil {
			fmt.Printf("error: %v\n", err)
   			 return
		}
		for _, scannedPort := range scannedPorts {
			if scannedPort.State == "Open" {
				openPorts = append(openPorts, scannedPort.Port)
			}
		}
		
		if output == "table" {
			fmt.Printf("open ports found: %d\n\n",len(openPorts))
			fmt.Printf("grabbing banners\n\n")
		}
		var serviceResults []ServiceResult
		for _, port := range openPorts {
			result, err := scanner.GrabBanner(target, port, time.Duration(timeout)*time.Second)
			if err != nil {
				fmt.Printf("port %d: %v\n", port, err)
				continue
			}
			if result.Banner == "" && result.Service == "" {
				if output == "table" {
					fmt.Printf("port %d: no banner\n\n", port)
				}
    			continue
			}
			serviceResults = append(serviceResults, ServiceResult{
				Port: result.Port,
				State: "Open",
				Banner: result.Banner,
				Service: result.Service,
				Version: result.Version,
				OS: result.OS,
			})
			if output == "table" {
				fmt.Println("results for port:", result.Port)
				fmt.Printf("raw banner: %v\n", result.Banner)
				fmt.Printf("service: %v\nversion: %v\nos :%v\n\n", result.Service, result.Version, result.OS)
			}
		}
		if output == "json" {
			err = utils.PrintJSON(serviceResults)
			if err != nil {
				fmt.Printf("error marshalling json: %v\n", err)
				return
			}
		}	
	},
}

func init() {
	serviceCmd.Flags().StringP("target", "t", "", "Target IP or hostname (required)")
	serviceCmd.Flags().Bool("all", false, "Port range 1-65535")
	serviceCmd.Flags().Bool("top100", false, "Use top 100 common ports")
	serviceCmd.Flags().Bool("top1000", false, "Use top 1000 common ports")

	serviceCmd.MarkFlagRequired("target")
}

type ServiceResult struct {
	Port    int    	`json:"port"`
    State   string 	`json:"state"`
    Banner  string 	`json:"banner,omitempty"`
    Service string 	`json:"service,omitempty"`
    Version string 	`json:"version,omitempty"`
    OS      string 	`json:"os,omitempty"`
}