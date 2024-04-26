package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	filelnfo, err := os.Stat("Nana")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if filelnfo.IsDir() {
		fmt.Println("Nana adalah sebuah direktori.")
	} else {
		fmt.Println("Nana adalah sebuah file.")
	}
}
