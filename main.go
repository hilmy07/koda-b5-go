package main

import (
	"fmt"
	"koda-b5-go/internals"
	"math"
)

func main() {
	var pilihan int

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Hitung Luas & Keliling Lingkaran")
		fmt.Println("2. Tampilkan Segitiga Siku")
		fmt.Println("3. Operasi Slice")
		fmt.Println("4. Tampilkan Struct Data Diri")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			luas, keliling := internals.HitungLuasdanKeliling(10)
			fmt.Println("Luas:", math.Round(luas))
			fmt.Println("Keliling:", math.Round(keliling))

		case 2:
			fmt.Println(internals.Segitigasiku(5))

		case 3:
			internals.Slice()

		case 4:
			internals.Struct()

		case 0:
			fmt.Println("Program selesai 👋")
			return // keluar dari main → program berhenti

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
