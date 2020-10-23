package csvplay

import (
	"io"
)

type CSVReader interface {
	Read() (record []string, err error)
}

type CSVWriter interface {
	Write(s []string) error
}

type CSVPlay struct {
	Input  CSVReader
	Output CSVWriter
}

func (c *CSVPlay) Order(order []int) error {
	return iterateRows(c.Input,
		func(row []string) error {
			var newRow []string
			for _, col := range order {
				newRow = append(newRow, row[col])
			}
			c.Output.Write(newRow)
			return nil
		})
}

func iterateRows(c CSVReader, action func([]string) error) error {
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
