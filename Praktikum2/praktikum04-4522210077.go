package main

import "fmt"

func main() {
	//deklarasi besarnya variabel panjang
    panjang := 5 
    //deklarasi besarnya variabel lebar
	lebar := 3
	//memanggil "hitungLuasKelilingPersegiPanjang" lalu hasil perhitungannya disimpan pada variabel luas dan keliling
    luas, keliling := hitungLuasKelilingPersegiPanjang(panjang, lebar) 
	//mencetak hasil perhitungan luas dan keliling
    fmt.Println("Luas persegi panjang:", luas)
    fmt.Println("Keliling persegi panjang:", keliling)
}

func hitungLuasKelilingPersegiPanjang(panjang int, lebar int) (luas int, keliling int) { //deklarasi fungsi
    //perhitungan luas persegi panjang
	luas = panjang * lebar
	//perhitungan keliling persegi panjang
    keliling = 2 * (panjang + lebar)
    return
}
