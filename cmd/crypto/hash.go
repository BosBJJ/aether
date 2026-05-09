package crypto

import (
	"fmt"

	"github.com/spf13/cobra"
	"aether/internal/crypto"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Generate hash of the text or a file",
	Long: `Generates a cryptographic hash of the input text or file.
You can use it with --input or --file and --algorithm can be SHA1, MD5 (both not recommended), SHA3-256, SHA3-512, SHA256, SHA512.`,

	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		input, _ := cmd.Flags().GetString("text")
		algorithm, _ := cmd.Flags().GetString("algorithm")
		
		if input == "" && file == "" {
			fmt.Println("Error: You must provide either --text or --file")
			cmd.Usage()
			return
		}
		
		if input != "" && file != "" {
			fmt.Println("Error: Use either --text or --file, not both.")
			return
		}

		var result string
		var err error

		if file != "" {
			result, err = crypto.HashFile(file, crypto.HashAlgorithm(algorithm))
		} else {
			result, err = crypto.Hash(input, crypto.HashAlgorithm(algorithm))
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		
		fmt.Printf("Algorithm %s: %s\n",algorithm, result)
	},
}

func init() {
	hashCmd.Flags().StringP("file", "f", "", "File to hash")
	hashCmd.Flags().StringP("text", "t", "", "Text to hash")
	hashCmd.Flags().StringP("algorithm", "a", "sha256", "Algorithm to use")
}