package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	var kategori string
	if usia < 18 {
		kategori = "anak-anak"
	} else {
		kategori = "dewasa"
	}

	fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", nama, kategori)
}