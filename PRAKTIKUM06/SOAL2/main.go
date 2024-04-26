package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	// Mengubah izin direktori.
	err = os.Chmod("Nana", 0077)
	if err != nil {
	fmt.Println("Error:", err)
	return
	}
	fmt.Println("Izin 'Nana' telah diubah menjadi 0077.")
	
}