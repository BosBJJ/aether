package scan


import (
	"fmt"
	"time"

	"aether/internal/scanner"

	"github.com/spf13/cobra"
)

var bannersCmd = &cobra.Command{
	Use:   "banners",
	Short: "Grabs banners from specified port",
	Long:  `Connects to a specified port on a target host and grabs the service banner. Supports HTTP, HTTPS, SSH, FTP, SMTP, IMAP, and POP3.`,

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		port, _ := cmd.Flags().GetInt("port")
		
		timeout, _ := cmd.Root().PersistentFlags().GetInt("timeout")

		result, err := scanner.GrabBanner(target, port, time.Duration(timeout)*time.Second)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Println("Results for port:", result.Port)
		fmt.Printf("Raw banner: %v\n", result.Banner)
		fmt.Printf("Service: %v\nVersion: %v\nOS :%v\n", result.Service, result.Version, result.OS)
	},
}

func init() {
	bannersCmd.Flags().StringP("target", "t", "", "Target IP or hostname (required)")
	bannersCmd.Flags().IntP("port", "p", 80, "Port to grab banner from")

	bannersCmd.MarkFlagRequired("target")
}