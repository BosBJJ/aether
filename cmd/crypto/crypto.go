package crypto

import "github.com/spf13/cobra"

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Cryptography utilities",
	Long:  `A powerful set of cryptographic tools including hashing, encryption, password management, and more.`,
}

func Init() {
	cryptoCmd.AddCommand(
		hashCmd,
		encryptCmd,
		decryptCmd,
		generateCmd,
		hmacCmd,
	)
}

func GetCryptoCmd() *cobra.Command {
	Init()
	return cryptoCmd
}