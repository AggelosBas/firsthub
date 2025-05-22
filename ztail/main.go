package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Println("Usage: go run . -c <number> <file1> [<file2> ...]")
		os.Exit(1)
	}

	count, err := strconv.Atoi(os.Args[2])
	if err != nil || count < 0 {
		fmt.Println("Invalid byte count")
		os.Exit(1)
	}

	files := os.Args[3:]
	hasError := false

	for i, fileName := range files {
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("open", fileName+":", err.Error())
			hasError = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", fileName)
		}

		if len(content) < count {
			fmt.Print(string(content))
		} else {
			fmt.Print(string(content[len(content)-count:]))
		}
	}

	if hasError {
		os.Exit(1)
	}
}
