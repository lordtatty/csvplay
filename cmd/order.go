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
		reader, err := getFileOrPipeReader(opener, cmd, args)
		if err != nil {
			return err
		}
		order := orderFlags.GetOrder()
		writer := &BufferedCSVWriter{
			Writer: csv.NewWriter(cmd.OutOrStdout()),
		}
		defer writer.Flush()
		csvplay := csvplay.CSVPlay{
			Input:  reader,
			Output: writer,
		}
		return csvplay.Order(order)
	}
}

func getFileOrPipeReader(opener CSVOpener, cmd *cobra.Command, args []string) (*csv.Reader, error) {
	r, err := getPipeCsvReader(cmd.InOrStdin())
	if err != nil {
		return nil, err
	}
	if r == nil {
		r, err = opener.Open(args[0])
		if err != nil {
			return nil, err
		}
	}
	return r, nil

}

type OrderFlags struct {
	F string
}

func (p *OrderFlags) GetOrder() []int {
	stringSlice := strings.Split(p.F, ",")
	order := make([]int, len(stringSlice))
	for i, s := range stringSlice {
		order[i], _ = strconv.Atoi(s)
		order[i]--
	}
	return order
}

var orderFlags OrderFlags

func init() {
	// rootCmd.AddCommand(rootCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&envConfigName, "env", "", "environment (eg. live / dev)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	OrderCmd.Flags().StringVar(&orderFlags.F, "f", "", "")
	// rootCmd.MarkFlagRequired("file")
}
