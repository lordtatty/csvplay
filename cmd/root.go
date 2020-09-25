package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// See https://godoc.org/github.com/pkg/errors for details on what is going on here
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				fmt.Printf("%+s:%d\n", f, f)
			}
		}
	}
}

// personaCmd represents the persona command
var rootCmd = &cobra.Command{
	Use:   "csvplay",
	Short: "Fun things with csv files",
	Long: `
	Lots of interesting info about CSV files
	`,
}

type CSVOpener interface {
	Open(filename string) (*csv.Reader, error)
}

type CSVFileOpener struct{}

func (c *CSVFileOpener) Open(filename string) (*csv.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(file), nil
}

type PipeFileChecker interface {
	Stat() (os.FileInfo, error)
}

func getPipeCsvReader(r io.Reader) (*csv.Reader, error) {
	p, ok := r.(PipeFileChecker)
	if !ok {
		return nil, nil
	}
	fi, err := p.Stat()
	if err != nil {
		return nil, err
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return nil, nil
	}
	r = bufio.NewReader(os.Stdin)
	return csv.NewReader(r), nil
}

func init() {
	// rootCmd.AddCommand(rootCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&envConfigName, "env", "", "environment (eg. live / dev)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(OrderCmd)
	rootCmd.AddCommand(JoinCmd)
	// rootCmd.MarkFlagRequired("file")
	// rootCmd.MarkFlagRequired("persona")
	// rootCmd.MarkFlagRequired("tenant")
}
