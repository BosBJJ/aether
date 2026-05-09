package crypto

import (
	"fmt"

	"github.com/spf13/cobra"
	"aether/internal/crypto"
)


var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates key or password",
	Long: `Will generate a key as a base64 or string password.
Use --key -k -length -l or --password -p -length -l -upper -u -special -s.
AES256 uses 32 byte keys`,

	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		key, _ := cmd.Flags().GetBool("key")
		password, _ := cmd.Flags().GetBool("password")
		special, _ := cmd.Flags().GetBool("special")
		upper, _ := cmd.Flags().GetBool("upper")
		
		if !key && !password {
			fmt.Println("Error: You must provide either --password -p or --key -k")
			cmd.Usage()
			return
		}
		
		if key && password {
			fmt.Println("Error: Use either --key or --password, not both.")
			return
		}
		if key && (special || upper) {
			fmt.Println("Error: --special and --upper only apply to passwords")
			return
		}


		if key {
			if length < 16 {
			fmt.Println("Error: --length must at least be 16 bytes")
			return
    		}
			newKey, err := crypto.GenerateKey(length)
			if err != nil {
				fmt.Printf("Error generating key: %v\n", err)
            	return
			}
			result := crypto.Base64Encode(newKey)
			fmt.Printf("Generated %v length key: %v\n", length, result)
		} else {
			if length < 12 {
			fmt.Println("Error: password --length has to be at least 12 characters")
			return
    		}
			pass, err := crypto.GeneratePassword(length, special, upper)
			if err != nil {
				fmt.Printf("Error generating password: %v\n", err)
            	return
			}
			fmt.Println("New password:", pass)
		}
	},
}

func init() {
	generateCmd.Flags().IntP("length", "l", 16, "Length of password or key")
	generateCmd.Flags().BoolP("key", "k", false, "Encryption Key")
	generateCmd.Flags().BoolP("password", "p", false, "Password")
	generateCmd.Flags().BoolP("special", "s", false, "Pass allows special characters")
	generateCmd.Flags().BoolP("upper", "u", false, "Pass has uppercase letters")
}