package cread

import (
	"bufio"
	"os"
	"strings"
)

func Readln () string {
	var reader = bufio.NewReader (os.Stdin)
	input, _ := reader.ReadString ('\n')
	// Remove final '\n' from input string.
	input = strings.Trim (input, "\n")

	return input
}

