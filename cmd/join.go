package cmd

import (
	"crypto/sha256"
	"encoding/csv"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var JoinCmd = &cobra.Command{
	Use:   "join",
	Short: "join columns together",
	Long: `
	Join some columns
		`,
	RunE: GetJoinCmd(&CSVFileOpener{}),
}

func GetJoinCmd(opener CSVOpener) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		csv1, err := getPipeCsvReader(cmd.InOrStdin())
		if err != nil {
			return err
		}
		if csv1 == nil {
			csv1, err = opener.Open(args[0])
			if err != nil {
				return err
			}
		} else {
			args = append([]string{""}, args...)
		}
		csv2, err := opener.Open(args[1])
		if err != nil {
			return err
		}
		stringSlice := strings.Split(JoinFlags.F, ":")
		key1, err := strconv.Atoi(stringSlice[0])
		if err != nil {
			return err
		}
		key2, err := strconv.Atoi(stringSlice[1])
		if err != nil {
			return err
		}
		key1--
		key2--

		csv2Cache, err := cacheByColumn(csv2, key2)
		if err != nil {
			return err
		}

		buffCSV := &BufferedCSVWriter{
			Writer: csv.NewWriter(cmd.OutOrStdout()),
		}
		defer buffCSV.Flush()
		sha256 := sha256.New()
		return iterateRows(csv1,
			func(row []string) error {
				hash := string(sha256.Sum([]byte(row[key1])))
				if _, ok := csv2Cache[hash]; ok {
					for k := range csv2Cache[hash] {
						newRow := append(row, csv2Cache[hash][k][1:]...)
						buffCSV.Write(newRow)
					}
				} else {
					buffCSV.Write(row)
				}
				return nil
			})
	}
}

func cacheByColumn(r *csv.Reader, col int) (map[string][][]string, error) {
	store := make(map[string][][]string)
	sha256 := sha256.New()
	err := iterateRows(r, func(s []string) error {
		if len(s) > col {
			hash := string(sha256.Sum([]byte(s[col])))
			if _, ok := store[hash]; !ok {
				store[hash] = [][]string{s}
			} else {
				store[hash] = append(store[hash], s)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return store, nil
}

var JoinFlags = struct {
	F string
}{}

func init() {
	JoinCmd.Flags().StringVar(&JoinFlags.F, "f", "", "")
}
