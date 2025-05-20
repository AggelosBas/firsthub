package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Println("Usage: go run . -c <number> <file1> [file2] ...")
		os.Exit(1)
	}

	n := 0
	_, err := fmt.Sscanf(os.Args[2], "%d", &n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid byte count")
		os.Exit(1)
	}

	files := os.Args[3:]
	exitCode := 0

	for i, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Printf("open %s: %v\n", name, err)
			exitCode = 1
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", name)
		}

		start := len(data) - n
		if start < 0 {
			start = 0
		}
		fmt.Printf("%s", data[start:])
	}

	os.Exit(exitCode)
}
