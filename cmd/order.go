package cmd

import (
	"encoding/csv"
	"strconv"
	"strings"

	"github.com/lordtatty/csvplay/csvplay"
	"github.com/spf13/cobra"
)

var OrderCmd = &cobra.Command{
	Use:   "order",
	Short: "Select columns",
	Long: `
	Select which colums you want to use
		`,
	RunE: GetOrderCmd(&CSVFileOpener{}),
}

func GetOrderCmd(opener CSVOpener) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		r, err := getPipeCsvReader(cmd.InOrStdin())
		if err != nil {
			return err
		}
		if r == nil {
			r, err = opener.Open(args[0])
			if err != nil {
				return err
			}
		}
		stringSlice := strings.Split(OrderFlags.F, ",")
		order := make([]int, len(stringSlice))
		for i, s := range stringSlice {
			order[i], _ = strconv.Atoi(s)
			order[i]--
		}
		buffCSV := &BufferedCSVWriter{
			Writer: csv.NewWriter(cmd.OutOrStdout()),
		}
		defer buffCSV.Flush()
		csvplay := csvplay.CSVPlay{
			Input:  r,
			Output: buffCSV,
		}
		return csvplay.Order(order)
	}
}

var OrderFlags = struct {
	F string
}{}

func init() {
	// rootCmd.AddCommand(rootCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&envConfigName, "env", "", "environment (eg. live / dev)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	OrderCmd.Flags().StringVar(&OrderFlags.F, "f", "", "")
	// rootCmd.MarkFlagRequired("file")
}
