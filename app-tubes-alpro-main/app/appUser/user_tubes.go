package user_tubes

import (
	"fmt"

	menu_tubes "github.com/Kidar12/tubes-alpro/app/appMenu"
	struct_tubes "github.com/Kidar12/tubes-alpro/app/appStruct"
)

func IsInput(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa *int) {
	/* input Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat dari Mahasiswa yang didaftarkan */
	var ProgramStud int
	i := *banyakMahasiswa
	fmt.Printf("Nama			      : ")
	fmt.Scan(&arr[i].NamaMahasiswa)

	fmt.Printf("Nomor Induk Kependudukan      : ")
	fmt.Scan(&arr[i].NomorInduk)

	fmt.Printf("Tanggal Lahir (Ex : 8 9 2004) : ")
	fmt.Scan(&arr[i].TanggalLahir, &arr[i].BulanLahir, &arr[i].TahunLahir)

	fmt.Print("Umur			      : ")
	fmt.Scan(&arr[i].UmurMahasiswa)

	fmt.Printf("Asal Kota 		      : ")
	fmt.Scan(&arr[i].KotaTinggal)

	fmt.Printf("Alamat Rumah		      : ")
	fmt.Scan(&arr[i].NamaJalan)

	fmt.Println()

	menu_tubes.MenuJurusan()
	fmt.Print("Pilih Jurusan		: ")
	fmt.Scan(&ProgramStud)
	switch ProgramStud {
	case 1:
		arr2[i].Program = "S1"
		arr2[i].Jurusan = "Informatika"
	case 2:
		arr2[i].Program = "S1"
		arr2[i].Jurusan = "Sistem Informasi"
	case 3:
		arr2[i].Program = "D4"
		arr2[i].Jurusan = "Informatika"
	case 4:
		arr2[i].Program = "D4"
		arr2[i].Jurusan = "Sistem Informasi"
	case 5:
		arr2[i].Program = "D3"
		arr2[i].Jurusan = "Informatika"
	case 6:
		arr2[i].Program = "D3"
		arr2[i].Jurusan = "Sistem Informasi"
	}
	fmt.Print("Masukan nilai anda	: ")
	fmt.Scan(&arr2[i].NilaiTes)
	i += 1
	fmt.Println()
	*banyakMahasiswa = i
}

func LihatStatusKelulusan(arr struct_tubes.TabArr, arr2 struct_tubes.TabArr2, banyakMahasiswa int, arr3 struct_tubes.TabKonf, CariNama string) {
	ada := false
	idx := -1
	for i := 0; i < banyakMahasiswa; i++ {
		if CariNama == arr[i].NamaMahasiswa {
			ada = true
			idx = i
		}
	}
	if ada == true {
		fmt.Printf(" Nama		  : %s		\n", arr[idx].NamaMahasiswa)
		fmt.Printf(" Nomor Induk	  : %s 			\n", arr[idx].NomorInduk)
		fmt.Printf(" Status Kelulusan : %t\n", arr3[idx].Konfirmasi)
	} else {
		fmt.Println("||	Nama tidak ditemukan.	||")
	}
}

func CekStatusMahasiswa(arr struct_tubes.TabArr, banyakMahasiswa int, CariNama string) {
	// Binary Search //
	idx := 0
	for i := 0; i < banyakMahasiswa; i++ {
		if arr[i].NamaMahasiswa == CariNama {
			idx += 1
		}
	}
	if idx != 0 {
		fmt.Println("||	Nama Sudah Terdaftar	||")
	} else {
		fmt.Println("||	Nama Belum Terdaftar	||")
	}
}
