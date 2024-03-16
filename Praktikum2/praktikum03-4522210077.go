package main

import "fmt" //library untuk menggunakan Println

func main() {
    panjang := 5 //deklarasi besarnya variabel panjang
    lebar := 3//deklarasi besarnya variabel lebar

    luas := hitungLuasPersegiPanjang(panjang, lebar) //deklarasi untuk menyimpan hasil perhitungan luas
    keliling := hitungKelilingPersegiPanjang(panjang, lebar)//deklarasi untuk menyimpan hasil perhitungan keliling

    fmt.Println("Luas persegi panjang:", luas) //mencetak nilai dari luas
    fmt.Println("Keliling persegi panjang:", keliling)//mencetak nilai dari Keliling
}

func hitungLuasPersegiPanjang(panjang int, lebar int) int {
    return panjang * lebar // mengembalikan hasil dari perkalian panjang dengan lebar
}

func hitungKelilingPersegiPanjang(panjang int, lebar int) int {
    return 2 * (panjang + lebar)// mengembalikan hasil dari perhitungan 2 dikali penjumlahan panjang dengan lebar
}
