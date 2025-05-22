package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		return
	}

	byteCount := 0
	for _, ch := range os.Args[2] {
		if ch < '0' || ch > '9' {
			return
		}
		byteCount = byteCount*10 + int(ch-'0')
	}

	files := os.Args[3:]
	errorOccurred := false

	for i, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			errorOccurred = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", file)
		}

		if byteCount > len(content) {
			fmt.Print(string(content))
		} else {
			fmt.Print(string(content[len(content)-byteCount:]))
		}
	}

	if errorOccurred {
		os.Exit(1)
	}
}
