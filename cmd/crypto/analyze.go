package crypto

import (
	"fmt"

	"github.com/spf13/cobra"
	"aether/internal/crypto"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze password strength",
	Long:  `Analyze the strength of a password and estimate how long it would take to crack.`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		if password == "" {
			fmt.Println("Error: Password cannot be empty")
			return
		}
        crypto.AnalyzePassword(password)
	},
}

func init() {
	
}