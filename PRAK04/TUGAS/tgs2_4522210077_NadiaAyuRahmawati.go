package main

import "fmt"

// Mengurutkan array dalam satu kali perulangan
func main() {
    fmt.Println("Bubble Short")
    var arrayNumber = [20]int{5, 3, 8, 1, 2, 7, 9, 4, 6} // Mengisi arrayNumber dengan contoh angka

    // Bubble sort
    for i := 0; i < len(arrayNumber); i++ {
        for j := 0; j < len(arrayNumber)-1-i; j++ {
            if arrayNumber[j] > arrayNumber[j+1] {
                // Tukar posisi jika elemen lebih besar dari elemen berikutnya
                arrayNumber[j], arrayNumber[j+1] = arrayNumber[j+1], arrayNumber[j]
            }
        }
    }

    fmt.Println("Setelah dilakukan pengurutan.")
    fmt.Println(arrayNumber)
}
