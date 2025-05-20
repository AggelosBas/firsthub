package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printRuneByRune(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printError(err error) {
	printRuneByRune("ERROR: ")
	printRuneByRune(err.Error())
	z01.PrintRune('\n')
}

func printFile(file *os.File) {
	buf := make([]byte, 1)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			z01.PrintRune(rune(buf[0]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			printError(err)
			break
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printFile(os.Stdin)
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		printFile(file)
		file.Close()
	}
}
