package crypto

import (
	"fmt"

	"aether/internal/crypto"
	"aether/internal/utils"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze password strength",
	Long:  `Analyze the strength of a password and estimate how long it would take to crack.`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		password := args[0]

        res := crypto.AnalyzePassword(password)

		if output == "json" {
			err := utils.PrintJSON(res)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
				return
			}
			return
		}

		fmt.Printf("Password Entropy: %.2f bits\n", res.Entropy)
		fmt.Printf("Crack Time : %s\n", res.CrackTime)
		for i, m := range res.MatchSequence {
			fmt.Printf(" %d. %-12s Token: %-15s Entropy: %.2f bits\n",
			i+1,
			"["+m.Pattern+"]",
			"'"+m.Token+"'",
			m.Entropy)
		}
	},
}

func init() {
	
}