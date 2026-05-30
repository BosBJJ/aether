package crypto

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/BosBJJ/aether/internal/crypto"
)

var hmacCmd = &cobra.Command{
	Use:   "hmac",
	Short: "Generate or verify hmac",
	Long: `Generates or verifies a  Hash-based Message Authentication Code.
You can use it with --input or --file and --algorithm can be SHA1, MD5 (both not recommended), SHA3-256, SHA3-512, SHA256, SHA512.`,

	Run: func(cmd *cobra.Command, args []string) {
		keyStr, _ := cmd.Flags().GetString("key")
		message, _ := cmd.Flags().GetString("message")
		algorithm, _ := cmd.Flags().GetString("algorithm")
		expected, _ := cmd.Flags().GetString("verify")
		
		if message == "" || keyStr == "" {
			fmt.Println("Error: You must provide --message and --key")
			cmd.Usage()
			return
		}

		var result string
		var err error
		if keyStr == "" {
			fmt.Println("Error: --key is required")
			return
    	}

		msg := []byte(message)
		key := []byte(keyStr)

		if expected != "" {
			valid, err := crypto.VerifyHMAC(msg, key, expected, crypto.HashAlgorithm(algorithm))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if valid {
				fmt.Println("HMAC Verification Successful")
			} else {
				fmt.Println("HMAC Verification FAILED, possible data tampering")
			}
			
		} else {
			result, err = crypto.GenerateHMAC(msg, key, crypto.HashAlgorithm(algorithm))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Algorithm %s: %s\n",algorithm, result)
		}
		
	},
}

func init() {
	hmacCmd.Flags().StringP("message", "m", "", "Message to hash")
	hmacCmd.Flags().StringP("key", "k", "", "Secret key")
	hmacCmd.Flags().StringP("algorithm", "a", "sha256", "Algorithm to use")
	hmacCmd.Flags().String("verify", "", "Expected HMAC to verify against")

}
