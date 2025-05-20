package main

import (
	"io"
	"os"
	"github.com/01-edu/z01"
)

func printErr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printFromReader(reader io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			for _, b := range buf[:n] {
				z01.PrintRune(rune(b))
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Χωρίς όρισμα: διάβασε από stdin
		printFromReader(os.Stdin)
	} else {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				printErr("ERROR: " + err.Error())
				os.Exit(1)
			}
			printFromReader(file)
			file.Close()
		}
	}
}
