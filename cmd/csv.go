package cmd

import (
	"encoding/csv"
	"io"
)

type BufferedCSVWriter struct {
	Writer  *csv.Writer
	buffLen int
}

func (b *BufferedCSVWriter) Write(s []string) error {
	err := b.Writer.Write(s)
	b.buffLen++
	if err != nil {
		return err
	}
	if b.buffLen > 1000 {
		b.Flush()
	}
	return nil
}

func (b *BufferedCSVWriter) Flush() {
	b.Writer.Flush()
	b.buffLen = 0
}

func iterateRows(c *csv.Reader, action func([]string) error) error {
	c.LazyQuotes = true
	for {
		row, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = action(row)
		if err != nil {
			return err
		}
	}
	return nil
}
