package main

import "fmt"

func main() {
	var menu, angka1, angka2, hasil int
	fmt.Println("Pilihan Operasi")
	fmt.Println("===============")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Riwayat")
	fmt.Println("9. EXIT")
	fmt.Print("Masukkan Pilihan :  ")
	fmt.Scanln(&menu)

}
