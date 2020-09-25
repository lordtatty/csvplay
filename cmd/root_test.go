package cmd_test

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type MockFileOpener struct {
	files map[string]string
	err   error
}

func (m *MockFileOpener) Open(filename string) (*csv.Reader, error) {
	if m.err != nil {
		return nil, m.err
	}
	var exists bool
	for k := range m.files {
		if k == filename {
			exists = true
		}
	}
	if exists != true {
		return nil, fmt.Errorf("Filename not found")
	}
	return csv.NewReader(strings.NewReader(m.files[filename])), nil
}
