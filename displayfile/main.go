package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:] // Παίρνουμε μόνο τα arguments (χωρίς το πρόγραμμα)

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	// Άνοιγμα αρχείου
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Αντιγραφή του περιεχομένου στην οθόνη (standard output)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
	}
}
