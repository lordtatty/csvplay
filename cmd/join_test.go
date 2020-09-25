package cmd_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/lordtatty/csvplay/cmd"
	"github.com/stretchr/testify/assert"
)

func TestJoinCLISettings(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("join", cmd.JoinCmd.Use)
	assert.Equal("join columns together", cmd.JoinCmd.Short)
	assert.Equal(`
	Join some columns
		`, cmd.JoinCmd.Long)
}

func TestJoin(t *testing.T) {
	csv1 := `"r1f1","r1f2","r1f3"
"r2f1","r2f2","r2f3"
"r3f1","r3f2","r3f3"`
	csv2 := `"r1f1","csv2r1f2","csv2r1f3"
"r2f1","csv2r2f2","csv2r2f3"
"r2f1","otherCsv2r2f2","otherCsv2r2f3"
"notLinked","csv2r3f2","csv2r3f3"`

	expected := `r1f1,r1f2,r1f3,csv2r1f2,csv2r1f3
r2f1,r2f2,r2f3,csv2r2f2,csv2r2f3
r2f1,r2f2,r2f3,otherCsv2r2f2,otherCsv2r2f3
r3f1,r3f2,r3f3
`
	opener := &MockFileOpener{
		files: map[string]string{
			"csv1.csv": csv1,
			"csv2.csv": csv2,
		},
	}
	b := bytes.NewBufferString("")

	cmd.JoinCmd.SetOutput(b)
	joinCmd := cmd.GetJoinCmd(opener)

	cmd.JoinFlags.F = "1:1"

	assert := assert.New(t)
	err := joinCmd(cmd.JoinCmd, []string{"csv1.csv", "csv2.csv"})
	assert.NoError(err)
	out, err := ioutil.ReadAll(b)
	assert.NoError(err)
	assert.Equal(expected, string(out))
}

func TestOrderJoinFileError(t *testing.T) {
	expectedErr := fmt.Errorf("Expected Error")
	opener := &MockFileOpener{
		err: expectedErr,
	}
	b := bytes.NewBufferString("")

	cmd.JoinCmd.SetOutput(b)
	joinCmd := cmd.GetJoinCmd(opener)

	cmd.JoinFlags.F = "1"

	assert := assert.New(t)
	err := joinCmd(cmd.JoinCmd, []string{"test.csv"})
	assert.EqualError(expectedErr, err.Error())
	out, err := ioutil.ReadAll(b)
	assert.NoError(err)
	assert.Equal("", string(out))
}
