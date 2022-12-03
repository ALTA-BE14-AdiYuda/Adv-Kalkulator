package main

import "fmt"

func main() {
	var menu, angka1, angka2, hasil int
	var riwayat []string
	for menu < 9 {
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
		if menu < 5 && menu > 0 {
			fmt.Print("Masukkan Angka Pertama :  ")
			fmt.Scanln(&angka1)
			fmt.Print("Masukkan Angka Kedua :  ")
			fmt.Scanln(&angka2)
			switch menu {
			case 1:
				hasil = angka1 + angka2
				riwayat = append(riwayat, fmt.Sprintf("%d + %d = %d\n", angka1, angka2, hasil))
			case 2:
				hasil = angka1 - angka2
				riwayat = append(riwayat, fmt.Sprintf("%d - %d = %d\n", angka1, angka2, hasil))
			case 3:
				hasil = angka1 * angka2
				riwayat = append(riwayat, fmt.Sprintf("%d * %d = %d\n", angka1, angka2, hasil))
			case 4:
				if angka2 == 0 {
					hasil = 0
				} else {
					hasil = angka1 / angka2
				}
				riwayat = append(riwayat, fmt.Sprintf("%d / %d = %d\n", angka1, angka2, hasil))
			}
		} else if menu == 5 {
			for i := 0; i < len
			(riwayat); i++ {
				fmt.Println(riwariwayat[i])
			}
		}
	}
}
