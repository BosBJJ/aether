package crypto

import (
	"fmt"

	"github.com/spf13/cobra"
	"aether/internal/crypto"
)


var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts file or text with AES256",
	Long: `Decrypts text or a file that was encrypted with AES-256-GCM.
Use --text -t or --file -f and --destination -d. Requires --key -k`,

	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		dest, _ := cmd.Flags().GetString("destination")
		input, _ := cmd.Flags().GetString("text")
		keyStr, _ := cmd.Flags().GetString("key")
		
		if input == "" && file == "" {
			fmt.Println("Error: You must provide either --text or --file")
			cmd.Usage()
			return
		}
		
		if input != "" && file != "" {
			fmt.Println("Error: Use either --text or --file, not both.")
			return
		}
		if keyStr == "" {
			fmt.Println("Error: --key is required")
			return
    	}

		key, _ := crypto.Base64Decode(keyStr)
		var err error

		if file != "" {
			if dest == "" {
            dest = file + ".aether"
        	}
			err = crypto.DecryptFileAES(file, dest, key)
			if err != nil {
				fmt.Printf("Error decrypting file: %v\n", err)
            	return
			}
		} else {
			plaintext, err := crypto.DecryptAES(input, key)
			if err != nil {
				fmt.Printf("Error decrypting text: %v\n", err)
            	return
			}
			fmt.Println("Decrypted:", string(plaintext))
		}
	},
}

func init() {
	decryptCmd.Flags().StringP("file", "f", "", "File to decrypt")
	decryptCmd.Flags().StringP("destination", "d", "", "Output file path")
	decryptCmd.Flags().StringP("text", "t", "", "Text to decrypt")
	decryptCmd.Flags().StringP("key", "k", "", "Encryption Key")
}