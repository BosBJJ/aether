package report

import (
	"github.com/BosBJJ/aether/cmd/config"
	"github.com/BosBJJ/aether/internal/database"
	"github.com/BosBJJ/aether/internal/utils"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved scans",
	Long:  `Displays a table of all saved scans with their ID, timestamp, command type and target.`,

	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		scans, err := database.ListScans(config.DB)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		
		if len(scans) == 0 {
			fmt.Println("No scans found.")
			return
		}

		if output == "json" {
			err = utils.PrintJSON(scans)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
			}
			return
		}

		fmt.Println("Saved Scans")
		fmt.Println("===============")
		fmt.Printf("%-6s %-25s %-15s %s\n","ID", "Timestamp", "Command Type", "Target")

		for _, scan := range scans {
			t, _ := time.Parse(time.RFC3339, scan.Timestamp)
			fmt.Printf("%-6d %-25s %-15s %s\n", scan.Id, t.Format("2006-01-02 15:04:05"), scan.CommandType, scan.Target)
		}
	},
}

func init() {
	
}