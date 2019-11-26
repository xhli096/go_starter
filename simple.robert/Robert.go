package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your name:")
	input, err := inputReader.ReadSlice('\n')
	if err != nil {
		fmt.Print("An error occurred: %s\n", err)
		return
	}
}
