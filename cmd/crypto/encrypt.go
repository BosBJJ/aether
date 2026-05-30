package crypto

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/BosBJJ/aether/internal/crypto"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts file or text with AES256",
	Long: `Encrypts a file to dest or encrypts text using AES.
You can use it with --text -t or --file -f and --destination -d. Requires --key -k`,

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
		plaintext := []byte(input)
		var err error

		if file != "" {
			if dest == "" {
            dest = file + ".aether"
        	}
			err = crypto.EncryptFileAES(file, dest, key)
			if err != nil {
				fmt.Printf("Error encrypting file: %v\n", err)
            	return
			}
		} else {
			encrypted, err := crypto.EncryptAES(plaintext, key)
			if err != nil {
				fmt.Printf("Error encrypting text: %v\n", err)
            	return
			}
			fmt.Println("Encrypted:", encrypted)
		}
	},
}

func init() {
	encryptCmd.Flags().StringP("file", "f", "", "File to encrypt")
	encryptCmd.Flags().StringP("destination", "d", "", "Output file path")
	encryptCmd.Flags().StringP("text", "t", "", "Text to encrypt")
	encryptCmd.Flags().StringP("key", "k", "", "Encryption Key")
}