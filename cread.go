package cread

import (
	"bufio"
	"fmt"
	"os"
)

func read (banner string) string {
	if banner != "" {
		fmt.Printf (banner)
	}
	var reader = bufio.NewReader (os.Stdin)
	text, _ := reader.ReadString ('\n')

	fmt.Print ("You typed %v.", text)

	return text
}

