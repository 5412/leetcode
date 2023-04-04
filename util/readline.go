package util

import (
	"bufio"
	"io"
)

func ReadEachline(r io.Reader, parser func(string) error) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parser(scanner.Text())
	}
	return nil
}
