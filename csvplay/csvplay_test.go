package csvplay_test

import (
	"io"
)

type CSVReaderMock struct {
	arr  [][]string
	curr int
}

func (c *CSVReaderMock) Read() (record []string, err error) {
	if c.curr+1 > len(c.arr) {
		return nil, io.EOF
	}
	i := c.curr
	c.curr++
	return c.arr[i], nil
}

type CSVWriterMock struct {
	arr [][]string
}

func (c *CSVWriterMock) Write(s []string) error {
	c.arr = append(c.arr, s)
	return nil
}

func (c *CSVWriterMock) Value() [][]string {
	return c.arr
}
