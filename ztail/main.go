package main

import (
	"fmt"
	"os"
)

func Atoi(s string) (int, bool) {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}

func main() {
	args := os.Args
	if len(args) < 4 || args[1] != "-c" {
		os.Exit(1)
	}

	count, ok := Atoi(args[2])
	if !ok {
		os.Exit(1)
	}

	files := args[3:]
	hadError := false

	for i, name := range files {
		file, err := os.Open(name)
		if err != nil {
			fmt.Printf("open %s: %s\n", name, err)
			hadError = true
			continue
		}

		info, err := file.Stat()
		if err != nil {
			fmt.Printf("open %s: %s\n", name, err)
			hadError = true
			file.Close()
			continue
		}

		size := info.Size()
		offset := int64(0)
		if size > int64(count) {
			offset = size - int64(count)
		}

		buf := make([]byte, count)
		n, err := file.ReadAt(buf, offset)
		file.Close()

		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", name)
		}
		fmt.Printf("%s", buf[:n])
	}

	if hadError {
		os.Exit(1)
	}
}
