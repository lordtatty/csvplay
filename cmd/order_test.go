package cmd_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/lordtatty/csvplay/cmd"
	"github.com/stretchr/testify/assert"
)

func TestOrderCLISettings(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("order", cmd.OrderCmd.Use)
	assert.Equal("Select columns", cmd.OrderCmd.Short)
	assert.Equal(`
	Select which colums you want to use
		`, cmd.OrderCmd.Long)
}

func TestOrder(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`
	expected := `first_name
Rob
Ken
Robert
`
	opener := &MockFileOpener{
		files: map[string]string{
			"test.csv": in,
		},
	}
	b := bytes.NewBufferString("")

	cmd.OrderCmd.SetOutput(b)
	orderCmd := cmd.GetOrderCmd(opener)

	cmd.OrderFlags.F = "1"

	assert := assert.New(t)
	err := orderCmd(cmd.OrderCmd, []string{"test.csv"})
	assert.NoError(err)
	out, err := ioutil.ReadAll(b)
	assert.NoError(err)
	assert.Equal(expected, string(out))
}

func TestOrderOpenFileError(t *testing.T) {
	expectedErr := fmt.Errorf("Expected Error")
	opener := &MockFileOpener{
		err: expectedErr,
	}
	b := bytes.NewBufferString("")

	cmd.OrderCmd.SetOutput(b)
	orderCmd := cmd.GetOrderCmd(opener)

	cmd.OrderFlags.F = "1"

	assert := assert.New(t)
	err := orderCmd(cmd.OrderCmd, []string{"test.csv"})
	assert.EqualError(expectedErr, err.Error())
	out, err := ioutil.ReadAll(b)
	assert.NoError(err)
	assert.Equal("", string(out))
}
