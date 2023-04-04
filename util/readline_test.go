package util

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadEachline(t *testing.T) {
	input := `
		string
		int
		float
		big int
	`
	ReadEachline(bytes.NewBuffer([]byte(input)), func(s string) error {
		fmt.Println(s)
		return nil
	})
}
