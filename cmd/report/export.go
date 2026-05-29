package report

import (
	"aether/cmd/config"
	"aether/internal/database"
	"aether/internal/utils"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the selected scan, use --output table, json, html",
	Long: `Exports a saved scan by ID in the specified format (table, json, html).
If --file is provided, writes the output to that file. Otherwise prints to stdout.

Examples:
  aether report export 42
  aether report export 42 -o html
  aether report export 42 -o html --file scan-42.html`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please input the id of the scan you want to export")
			return
		}
		output, _ := cmd.Root().PersistentFlags().GetString("output")
		dst, _ := cmd.Flags().GetString("file")
		reportID, _ := strconv.Atoi(args[0])

		report, err := database.GetScan(config.DB, reportID)
		if err != nil {
			fmt.Printf("error retrieving report: %v\n", err)
			return
		}
		formatted, err := utils.FormatScan(report, output)
		if err != nil {
			fmt.Printf("error formatting report: %v\n", err)
			return
		}

		if dst == "" {
			fmt.Println(formatted)
			return
		}

		ext := output
		if ext == "table" {
			ext = "md"
		}
		if filepath.Ext(dst) == "" {
			dst = dst + "." + ext
		}



		err = os.WriteFile(dst, []byte(formatted), 0644)
		if err != nil {
			fmt.Printf("error creating file %v: %v", dst, err)
			return
		}

		fmt.Printf("Exported report %v to %v\n", reportID, dst)

	},
}

func init() {
	exportCmd.Flags().StringP("file", "f", "", "file destination")
}
