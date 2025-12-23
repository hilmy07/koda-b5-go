package internals

import "fmt"

func Struct() {
	// struct untuk riwayat pendidikan
	type Pendidikan struct {
		Nama    string
		Jurusan string
	}

	// struct utama data diri
	type DataDiri struct {
		Nama              string
		Foto              string
		Email             string
		Umur              int
		NomorTelepon      string
		StatusPernikahan  string
		RiwayatPendidikan Pendidikan
	}

	data := DataDiri{
		Nama:             "Muhammad Hilmy",
		Foto:             "foto.jpg",
		Email:            "hilmy@mail.com",
		Umur:             22,
		NomorTelepon:     "08123456789",
		StatusPernikahan: "Belum Menikah",
		RiwayatPendidikan: Pendidikan{
			Nama:    "Universitas XYZ",
			Jurusan: "Informatika",
		},
	}

	fmt.Println("Nama:", data.Nama)
	fmt.Println("Foto:", data.Foto)
	fmt.Println("Email:", data.Email)
	fmt.Println("Umur:", data.Umur)
	fmt.Println("No Telp:", data.NomorTelepon)
	fmt.Println("Status Pernikahan:", data.StatusPernikahan)
	fmt.Println("Pendidikan:", data.RiwayatPendidikan.Nama)
	fmt.Println("Jurusan:", data.RiwayatPendidikan.Jurusan)
}