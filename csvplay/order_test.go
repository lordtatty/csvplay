package csvplay_test

import (
	"testing"

	"github.com/lordtatty/csvplay/csvplay"
	"github.com/stretchr/testify/assert"
)

func TestOrderHappyPath(t *testing.T) {
	in := [][]string{
		{"first_name", "last_name", "username"},
		{"rob", "pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	tests := []struct {
		Name        string
		Order       []int
		ExpectedOut [][]string
	}{
		{
			Name:  "Select First Column",
			Order: []int{0},
			ExpectedOut: [][]string{
				{"first_name"},
				{"rob"},
				{"Ken"},
				{"Robert"},
			},
		},
		{
			Name:  "Select Second Column",
			Order: []int{1},
			ExpectedOut: [][]string{
				{"last_name"},
				{"pike"},
				{"Thompson"},
				{"Griesemer"},
			},
		},
		{
			Name:  "Select Third Column",
			Order: []int{2},
			ExpectedOut: [][]string{
				{"username"},
				{"rob"},
				{"ken"},
				{"gri"},
			},
		},
		{
			Name:  "Select Third and First Column",
			Order: []int{2, 0},
			ExpectedOut: [][]string{
				{"username", "first_name"},
				{"rob", "rob"},
				{"ken", "Ken"},
				{"gri", "Robert"},
			},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)
			in := &CSVReaderMock{
				arr: in,
			}
			out := &CSVWriterMock{}
			sut := csvplay.CSVPlay{
				Input:  in,
				Output: out,
			}
			sut.Order(tc.Order)
			assert.Equal(tc.ExpectedOut, out.Value())
		})
	}
}
