package report

import (
	"aether/cmd/config"
	"aether/internal/database"
	"aether/internal/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the selected scan",
	Long:  `Returns results for selected scan, including the type of command used, the results in raw form and a brief summary.`,

	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Root().PersistentFlags().GetString("output")

		if len(args) == 0 {
			fmt.Println("please input the id of the scan you want to see")
			return
		}
		reportID, _ := strconv.Atoi(args[0])

		scan, err := database.GetScan(config.DB, reportID)
		if err != nil {
			fmt.Printf("error fetching report: %v\n", err)
			return
		}

		if output == "json" {
			err = utils.PrintJSON(scan)
			if err != nil {
				fmt.Printf("error printing json: %v\n", err)
			}
			return
		}

		fmt.Println("Scan Results")
		fmt.Println("=================")
		fmt.Printf("%-25s: %-15v\n", "Target", scan.Target)
		fmt.Printf("%-25s: %-15v\n", "Command Type", scan.CommandType)
		fmt.Printf("%-25s: %-15v\n", "Time Scanned", scan.Timestamp)
		fmt.Printf("%-25s\n", "Summary")
		fmt.Println("==================")
		fmt.Printf("%-15v\n", scan.Summary)
		fmt.Println()
		fmt.Printf("%-25s\n", "Raw Result")
		fmt.Println("==================")
		fmt.Printf("%-15v\n", scan.RawResult)

	},
}

func init() {

}
