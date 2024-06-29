package sorting_tubes

import (
	"fmt"
	struct_tubes "github.com/Kidar12/tubes-alpro/app/appStruct"
	menu_tubes "github.com/Kidar12/tubes-alpro/app/appMenu"
)

func NilaiAscending(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		idx := pass - 1
		for j := pass; j < banyakMahasiswa; j++ {
			if (*arr2)[idx].NilaiTes > (*arr2)[j].NilaiTes {
				idx = j
			}
		}
		temp := (*arr2)[idx].NilaiTes
		temp1 := (*arr)[idx].NamaMahasiswa
		temp2 := (*arr)[idx].NomorInduk

		(*arr2)[idx].NilaiTes = (*arr2)[pass - 1].NilaiTes
		(*arr)[idx].NamaMahasiswa = (*arr)[pass - 1].NamaMahasiswa
		(*arr)[idx].NomorInduk = (*arr)[pass - 1].NomorInduk

		(*arr2)[pass - 1].NilaiTes = temp
		(*arr)[pass - 1].NamaMahasiswa = temp1
		(*arr)[pass - 1].NomorInduk = temp2
	}

	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		fmt.Printf("Nama	 : %s\nNo.Induk : %s\nNilai	 : %0.2f\n", arr[k].NamaMahasiswa, arr[k].NomorInduk, arr2[k].NilaiTes)
		fmt.Println("========================================")
	}
}

func NilaiDescending(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		j := pass - 1
		temp := (*arr2)[pass].NilaiTes
		temp1 := (*arr)[pass].NamaMahasiswa
		temp2 := (*arr)[pass].NomorInduk
		for j >= 0 && temp > (*arr2)[j].NilaiTes {
			(*arr)[j+1].NamaMahasiswa = (*arr)[j].NamaMahasiswa
			(*arr)[j+1].NomorInduk = (*arr)[j].NomorInduk
			(*arr2)[j+1].NilaiTes = (*arr2)[j].NilaiTes
			j--
		}
		(*arr2)[j+1].NilaiTes = temp
		(*arr)[j+1].NamaMahasiswa = temp1
		(*arr)[j+1].NomorInduk = temp2
	}

	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		fmt.Printf("Nama	 : %s\nNo.Induk : %s\nNilai	 : %0.2f\n", arr[k].NamaMahasiswa, arr[k].NomorInduk, arr2[k].NilaiTes)
		fmt.Println("========================================")
	}
}

func SortingJurusan(arr struct_tubes.TabArr, arr2 struct_tubes.TabArr2, banyakMahasiswa int) {
	var sortingJurusan int
	menu_tubes.SortingJurusan()
	fmt.Print("Pilih Menu Program dan Jurusan: ")
	fmt.Scan(&sortingJurusan)
	fmt.Println()
	switch sortingJurusan {
	case 1:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "S1" && arr2[k].Jurusan == "Informatika" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: S1 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}
		
	case 2:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "S1" && arr2[k].Jurusan == "Sistem Informasi" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: S1 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}
		
	case 3:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D4" && arr2[k].Jurusan == "Informatika" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D4 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}
		
	case 4:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D4" && arr2[k].Jurusan == "Sistem Informasi" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D4 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}
		
	case 5:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D3" && arr2[k].Jurusan == "Informatika" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D3 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}
		
	case 6:
		fmt.Println("========================================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr2[k].Program == "D3" && arr2[k].Jurusan == "Sistem Informasi" {
				fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D3 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println("========================================")
			}
		}
		
	}
}

func SortingLulus(arr struct_tubes.TabArr, arr3 struct_tubes.TabKonf, banyakMahasiswa int) {
	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		if arr3[k].Konfirmasi == true {
			fmt.Printf("Nama	 : %s\nNo.Induk : %s\nStatus	 : Lulus\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
			fmt.Println("========================================")
		}
	}
}

func SortingDitolak(arr struct_tubes.TabArr, arr3 struct_tubes.TabKonf, banyakMahasiswa int) {
	fmt.Println("========================================")
	for k := 0; k < banyakMahasiswa; k++ {
		if arr3[k].Konfirmasi == false {
			fmt.Printf("Nama	 : %s\nNo.Induk : %s\nStatus	 : Tidak Lulus\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
			fmt.Println("========================================")
		}
	}
}

func SortingUmur(arr *struct_tubes.TabArr, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		j := pass - 1
		temp := (*arr)[pass].UmurMahasiswa
		temp1 := (*arr)[pass].NamaMahasiswa
		temp2 := (*arr)[pass].NomorInduk
		for j >= 0 && temp > (*arr)[j].UmurMahasiswa {
			(*arr)[j+1].NamaMahasiswa = (*arr)[j].NamaMahasiswa
			(*arr)[j+1].NomorInduk = (*arr)[j].NomorInduk
			(*arr)[j+1].UmurMahasiswa = (*arr)[j].UmurMahasiswa
			j--
		}
		(*arr)[j+1].UmurMahasiswa = temp
		(*arr)[j+1].NamaMahasiswa = temp1
		(*arr)[j+1].NomorInduk = temp2
	}
	var umur int
	fmt.Print("Masukan Umur yang akan dicari: ")
	fmt.Scan(&umur)
	fmt.Println()
	ki := 0
	ka := banyakMahasiswa - 1
	idx := -1
	for ki <= ka && idx == -1 {
		te := (ka + ki)/2
		if arr[te].UmurMahasiswa < umur {
			ka = te - 1
		} else if arr[te].UmurMahasiswa > umur {
			ki = te + 1
		} else if arr[te].UmurMahasiswa == umur {
			idx = te
		}
	}
	if idx >= 0 {
		fmt.Println("========================================")
		for i := 0; i < banyakMahasiswa; i++ {
			if arr[i].UmurMahasiswa == umur {
				fmt.Printf("Nama	 : %s\nNo.Induk : %s\numur	 : %d\n", arr[i].NamaMahasiswa, arr[i].NomorInduk, arr[i].UmurMahasiswa)
				fmt.Println("========================================")
			}
		}
	} else {
		fmt.Printf("Tidak Terdapat Mahasiswa Dengan Umur %d\n", umur)
	}
}

/*
2
1
Shadra_Mehdi_Irfani
213123
8 9 2004
18
bandung
asda
1
90
2
1
Rizkydarmawan
2131232
11 9 2005
19
Bekasi
indah
1
88
2
1
shadrairfani
21312321
8 9 2004
19
bandung
asda
1
90
*/