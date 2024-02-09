package main

import (
	"fmt"
	"lol/internal"
	"os"
)

func main() {
	lines, err := internal.GetFileContents(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if lines == nil {
		return
	}

	program, err := internal.TokenizeProgram(&lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	internal.Interpret(&program)
}
