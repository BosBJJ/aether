package report

import (
	"aether/cmd/config"
	"aether/internal/database"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the selected scan",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please input the id of the scan you want to delete")
			return
		}
		reportID, _ := strconv.Atoi(args[0])

		err := database.DeleteScan(config.DB, reportID)
		if err != nil {
			fmt.Printf("error deleting report: %v\n", err)
			return
		}

		fmt.Printf("Deleted report %v\n", reportID)

	},
}

func init() {

}
