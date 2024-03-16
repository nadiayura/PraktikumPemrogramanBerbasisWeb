package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	mahasiswas := make(map[string]Mahasiswa)

	// Menambahkan data mahasiswa ke dalam map
	mahasiswas["4522109"] = Mahasiswa{"Ramadhani", "4522109", "Informatika"}
	mahasiswas["4622177"] = Mahasiswa{"Nadia Ayu", "4622177", "Teknik Elektro"}
	mahasiswas["4722166"] = Mahasiswa{"Salwa", "4722166", "Teknik Elektro"}
	mahasiswas["4822142"] = Mahasiswa{"Farah", "4822142", "Teknik Elektro"}
	mahasiswas["4922144"] = Mahasiswa{"Daffa", "4922144", "Teknik Elektro"}

	// Menampilkan daftar nama mahasiswa dengan perulangan
	fmt.Println("Daftar Mahasiswa:")
	for _, mhs := range mahasiswas {
		fmt.Println("-------------------------------")
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("NPM:", mhs.NPM)
		fmt.Println("Jurusan:", mhs.Jurusan)
		fmt.Println("-------------------------------")
	}

	// Mencari data mahasiswa berdasarkan NPM
	var npmCari string
	fmt.Print("\nMasukkan NPM mahasiswa yang ingin dicari: ")
	fmt.Scanln(&npmCari)
	fmt.Println("-------------------------------")

	mhs, found := mahasiswas[npmCari]
	if found {
		fmt.Println("\nMahasiswa dengan NPM", npmCari, "ditemukan.")
		fmt.Println("===============================")
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("NPM:", mhs.NPM)
		fmt.Println("Jurusan:", mhs.Jurusan)
		fmt.Println("===============================")
	} else {
		fmt.Println("\nMahasiswa dengan NPM", npmCari, "tidak ditemukan.")
	}
}
