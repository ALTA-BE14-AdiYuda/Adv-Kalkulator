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
	fmt.Print("Masukkan Angka Pertama :  ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan Angka Kedua :  ")
	fmt.Scanln(&angka2)
	switch menu {
	case 1:
		hasil = angka1 + angka2
	case 2:
		hasil = angka1 - angka2
	case 3:
		hasil = angka1 * angka2
	case 4:
		if angka2 == 0 {
			hasil = 0
		} else {
			hasil = angka1 / angka2
		}
	}

}
