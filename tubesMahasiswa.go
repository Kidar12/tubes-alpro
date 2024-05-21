package main

import (
	"fmt"
	"os"
	"os/exec"
)

type tabPenerimaan struct {
	namaMahasiswa1, namaMahasiswa2, namaMahasiswa3, nomorInduk string
	kotaTinggal, namaJalan1, namajalan2                        string
	tanggalLahir, bulanLahir, tahunLahir, umurMahasiswa        int
}

type tabRegis struct {
	program, jurusan string
	nilaiTes                       float64
}

const NMAX int = 1000

type tabArr [NMAX]tabPenerimaan
type tabArr2 [NMAX]tabRegis

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	fmt.Println("=====Form Pendaftaran Mahasiswa Baru Telkom University====")
	fmt.Println("=====A). Isi Data Diri Mahasiswa==========================")
	fmt.Printf("||     1. Daftar Mahasiswa				||\n")
	fmt.Printf("||     2. Edit Daftar Mahasiswa				||\n")
	fmt.Printf("||     3. Lihat Daftar Mahasiswa			||\n")
	fmt.Printf("||     4. Jenis Keluaran				||\n")
	fmt.Printf("||     5. Data Informasi Mahasiswa			||\n")
	fmt.Printf("||     6. Exit						||\n")
	fmt.Printf("==========================================================\n")
	fmt.Print("Pilih menu: ")
}

func menu2() {
	fmt.Println("=====B). Data Registrasi=================================")
	fmt.Printf("||     1. Daftar Seleksi				||\n")
	fmt.Printf("||     2. Riwayat Registrasi				||\n")
	fmt.Printf("||     3. Exit						||\n")
	fmt.Printf("==========================================================\n")
	fmt.Print("Pilih menu: ")
}

func menuJurusan() {
	fmt.Println("==================================")
	fmt.Println("|| 1. S1 Informatika		||")
	fmt.Println("|| 2. S1 Sistem Informasi	||")
	fmt.Println("|| 3. D4 Informatika		||")
	fmt.Println("|| 4. D4 Sistem Informasi 	||")
	fmt.Println("|| 5. D3 Informatika		||")
	fmt.Println("|| 6. D3 Sistem Informasi	||")
	fmt.Println("==================================")
}

func main() {
	var arr tabArr
	var arr2 tabArr2
	var banyakMahasiswa, chooseMenu, sukuKata int
	var lanjut bool = true

	banyakMahasiswa = 0
	for lanjut {
		menu()
		fmt.Scan(&chooseMenu)
		fmt.Println()
		if chooseMenu > 5 {
			lanjut = false
		}
		switch chooseMenu {
		case 1:
			isInput(&arr, &arr2, &banyakMahasiswa, sukuKata)
		case 2:
			isEdit(&arr, &arr2, &banyakMahasiswa, sukuKata)
		case 3:
			isOutput(arr, arr2, banyakMahasiswa)
		case 4:
			jenisOutput(arr, arr2, banyakMahasiswa)
		case 5:
			oldestYoungestAge(arr, banyakMahasiswa)
		}
		fmt.Println()
	}
}

func isInput(A *tabArr, B *tabArr2, n *int, sukuKata int) {
	/* input Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat dari Mahasiswa yang didaftarkan */
	var ProgramStud int
	i := *n
	fmt.Print("Berapakah Suku Kata Nama Anda? (1/2/3): ")
	fmt.Scan(&sukuKata)
	switch sukuKata {
	case 1:
		fmt.Printf("Nama			     : ")
		fmt.Scan(&A[i].namaMahasiswa1)

		fmt.Printf("Nomor Induk Kependudukan     : ")
		fmt.Scan(&A[i].nomorInduk)

		fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
		fmt.Scan(&A[i].tanggalLahir, &A[i].bulanLahir, &A[i].tahunLahir)

		fmt.Print("Umur			     : ")
		fmt.Scan(&A[i].umurMahasiswa)

		fmt.Printf("Asal Kota (Ex : Bandung)     : ")
		fmt.Scan(&A[i].kotaTinggal)

		fmt.Printf("Alamat Rumah (2 suku kata)   : ")
		fmt.Scan(&A[i].namaJalan1, &A[i].namajalan2)

		fmt.Println()

		for {
			menu2()
			var chooseMenu2 int
			fmt.Scan(&chooseMenu2)
			fmt.Println()
			if chooseMenu2 > 2 {
				break
			}
			switch chooseMenu2 {
			case 1:
				menuJurusan()
				fmt.Print("Pilih jurusan: ")
				fmt.Scan(&ProgramStud)
				switch ProgramStud {
				case 1:
					B[i].program = "S1"
					B[i].jurusan = "Informatika"
				case 2:
					B[i].program = "S1"
					B[i].jurusan = "Sistem Informasi"
				case 3:
					B[i].program = "D4"
					B[i].jurusan = "Informatika"
				case 4:
					B[i].program = "D4"
					B[i].jurusan = "Sistem Informasi"
				case 5:
					B[i].program = "D3"
					B[i].jurusan = "Informatika"
				case 6:
					B[i].program = "D3"
					B[i].jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&B[i].nilaiTes)
				fmt.Println()
			case 2:
				fmt.Printf("Rekap data nilai: \n")
				fmt.Printf("Program dan Jurusan Yang Anda Pilih: %s %s\n", B[i].program, B[i].jurusan)
				fmt.Printf("Nilai Yang Anda Peroleh: %0.1f\n", B[i].nilaiTes)
				fmt.Println()
			}
		}
		i += 1
		fmt.Println()
	case 2:
		fmt.Printf("Nama			     : ")
		fmt.Scan(&A[i].namaMahasiswa1, &A[i].namaMahasiswa2)

		fmt.Printf("Nomor Induk Kependudukan     : ")
		fmt.Scan(&A[i].nomorInduk)

		fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
		fmt.Scan(&A[i].tanggalLahir, &A[i].bulanLahir, &A[i].tahunLahir)

		fmt.Print("Umur			     : ")
		fmt.Scan(&A[i].umurMahasiswa)

		fmt.Printf("Asal Kota (Ex : Bandung)     : ")
		fmt.Scan(&A[i].kotaTinggal)

		fmt.Printf("Alamat Rumah (2 suku kata)   : ")
		fmt.Scan(&A[i].namaJalan1, &A[i].namajalan2)

		fmt.Println()

		for {
			menu2()
			var chooseMenu2 int
			fmt.Scan(&chooseMenu2)
			fmt.Println()
			if chooseMenu2 > 2 {
				break
			}
			switch chooseMenu2 {
			case 1:
				menuJurusan()
				fmt.Print("Pilih jurusan: ")
				fmt.Scan(&ProgramStud)
				switch ProgramStud {
				case 1:
					B[i].program = "S1"
					B[i].jurusan = "Informatika"
				case 2:
					B[i].program = "S1"
					B[i].jurusan = "Sistem Informasi"
				case 3:
					B[i].program = "D4"
					B[i].jurusan = "Informatika"
				case 4:
					B[i].program = "D4"
					B[i].jurusan = "Sistem Informasi"
				case 5:
					B[i].program = "D3"
					B[i].jurusan = "Informatika"
				case 6:
					B[i].program = "D3"
					B[i].jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&B[i].nilaiTes)
				fmt.Println()
			case 2:
				fmt.Printf("Rekap data nilai: \n")
				fmt.Printf("Program: %s %s\n", B[i].program, B[i].jurusan)
				fmt.Printf("Nilai: %0.1f\n", B[i].nilaiTes)
				fmt.Println()
			}
		}
		i += 1
		fmt.Println()
	case 3:
		fmt.Printf("Nama			     : ")
		fmt.Scan(&A[i].namaMahasiswa1, &A[i].namaMahasiswa2, &A[i].namaMahasiswa3)

		fmt.Printf("Nomor Induk Kependudukan     : ")
		fmt.Scan(&A[i].nomorInduk)

		fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
		fmt.Scan(&A[i].tanggalLahir, &A[i].bulanLahir, &A[i].tahunLahir)

		fmt.Print("Umur			     : ")
		fmt.Scan(&A[i].umurMahasiswa)

		fmt.Printf("Asal Kota (Ex : Bandung)     : ")
		fmt.Scan(&A[i].kotaTinggal)

		fmt.Printf("Alamat Rumah (2 suku kata)   : ")
		fmt.Scan(&A[i].namaJalan1, &A[i].namajalan2)

		fmt.Println()

		for {
			menu2()
			var chooseMenu2 int
			fmt.Scan(&chooseMenu2)
			fmt.Println()
			if chooseMenu2 > 2 {
				break
			}
			switch chooseMenu2 {
			case 1:
				menuJurusan()
				fmt.Print("Pilih jurusan: ")
				fmt.Scan(&ProgramStud)
				switch ProgramStud {
				case 1:
					B[i].program = "S1"
					B[i].jurusan = "Informatika"
				case 2:
					B[i].program = "S1"
					B[i].jurusan = "Sistem Informasi"
				case 3:
					B[i].program = "D4"
					B[i].jurusan = "Informatika"
				case 4:
					B[i].program = "D4"
					B[i].jurusan = "Sistem Informasi"
				case 5:
					B[i].program = "D3"
					B[i].jurusan = "Informatika"
				case 6:
					B[i].program = "D3"
					B[i].jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&B[i].nilaiTes)
				fmt.Println()
			case 2:
				fmt.Printf("Rekap data nilai: \n")
				fmt.Printf("Program: %s %s\n", B[i].program, B[i].jurusan)
				fmt.Printf("Nilai: %0.1f\n", B[i].nilaiTes)
				fmt.Println()
			}
		}
		i += 1
		fmt.Println()
	}
	var clear string
	fmt.Print("Bersihkan layar (yes/no): ")
	fmt.Scan(&clear)
	if clear == "yes" {
		clearScreen()
	}
	*n = i
}

func isEdit(A *tabArr, B *tabArr2, n *int, sukuKata int) {
	/* Mengedit mahasiswa yang sudah terdaftar dengan memasukan Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat yang baru */
	var suku int
	var menuEdit string
	var edit1, edit2, edit3, nomorIndukEdit, kotaTinggalEdit string
	var namaJalan1Edit, namajalan2Edit string
	var tanggalLahirEdit, bulanLahirEdit, tahunLahirEdit, umurMahasiswaEdit int
	var noIndukHapus string

	var programStud int
	var nilaiTesEdit float64

	fmt.Print("Menu: (tambah / edit / hapus): ")
	fmt.Scan(&menuEdit)

	switch menuEdit {
	case "tambah":
		i := *n
		for {
			fmt.Print("Berapakah Suku Kata Nama Anda? (1/2/3): ")
			fmt.Scan(&sukuKata)
			if sukuKata == 0 {
				break
			}
			switch sukuKata {
			case 1:
				fmt.Printf("Nama			     : ")
				fmt.Scan(&A[i].namaMahasiswa1)

				fmt.Printf("Nomor Induk Kependudukan     : ")
				fmt.Scan(&A[i].nomorInduk)

				fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
				fmt.Scan(&A[i].tanggalLahir, &A[i].bulanLahir, &A[i].tahunLahir)

				fmt.Print("Umur			     : ")
				fmt.Scan(&A[i].umurMahasiswa)

				fmt.Printf("Asal Kota (Ex : Bandung)     : ")
				fmt.Scan(&A[i].kotaTinggal)

				fmt.Printf("Alamat Rumah (2 suku kata)   : ")
				fmt.Scan(&A[i].namaJalan1, &A[i].namajalan2)

				fmt.Println()

				menuJurusan()
				fmt.Print("Pilih jurusan: ")
				fmt.Scan(&programStud)
				switch programStud {
				case 1:
					B[i].program = "S1"
					B[i].jurusan = "Informatika"
				case 2:
					B[i].program = "S1"
					B[i].jurusan = "Sistem Informasi"
				case 3:
					B[i].program = "D4"
					B[i].jurusan = "Informatika"
				case 4:
					B[i].program = "D4"
					B[i].jurusan = "Sistem Informasi"
				case 5:
					B[i].program = "D3"
					B[i].jurusan = "Informatika"
				case 6:
					B[i].program = "D3"
					B[i].jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&B[i].nilaiTes)
				fmt.Println()

				i += 1
				fmt.Println()
			case 2:
				fmt.Printf("Nama			     : ")
				fmt.Scan(&A[i].namaMahasiswa1, &A[i].namaMahasiswa2)

				fmt.Printf("Nomor Induk Kependudukan     : ")
				fmt.Scan(&A[i].nomorInduk)

				fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
				fmt.Scan(&A[i].tanggalLahir, &A[i].bulanLahir, &A[i].tahunLahir)

				fmt.Print("Umur			     : ")
				fmt.Scan(&A[i].umurMahasiswa)

				fmt.Printf("Asal Kota (Ex : Bandung)     : ")
				fmt.Scan(&A[i].kotaTinggal)

				fmt.Printf("Alamat Rumah (2 suku kata)   : ")
				fmt.Scan(&A[i].namaJalan1, &A[i].namajalan2)

				fmt.Println()

				menuJurusan()
				fmt.Print("Pilih jurusan: ")
				fmt.Scan(&programStud)
				switch programStud {
				case 1:
					B[i].program = "S1"
					B[i].jurusan = "Informatika"
				case 2:
					B[i].program = "S1"
					B[i].jurusan = "Sistem Informasi"
				case 3:
					B[i].program = "D4"
					B[i].jurusan = "Informatika"
				case 4:
					B[i].program = "D4"
					B[i].jurusan = "Sistem Informasi"
				case 5:
					B[i].program = "D3"
					B[i].jurusan = "Informatika"
				case 6:
					B[i].program = "D3"
					B[i].jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&B[i].nilaiTes)
				fmt.Println()

				i += 1
				fmt.Println()
			case 3:
				fmt.Printf("Nama			     : ")
				fmt.Scan(&A[i].namaMahasiswa1, &A[i].namaMahasiswa2, &A[i].namaMahasiswa3)

				fmt.Printf("Nomor Induk Kependudukan     : ")
				fmt.Scan(&A[i].nomorInduk)

				fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
				fmt.Scan(&A[i].tanggalLahir, &A[i].bulanLahir, &A[i].tahunLahir)

				fmt.Print("Umur			     : ")
				fmt.Scan(&A[i].umurMahasiswa)

				fmt.Printf("Asal Kota (Ex : Bandung)     : ")
				fmt.Scan(&A[i].kotaTinggal)

				fmt.Printf("Alamat Rumah (2 suku kata)   : ")
				fmt.Scan(&A[i].namaJalan1, &A[i].namajalan2)

				fmt.Println()

				menuJurusan()
				fmt.Print("Pilih jurusan: ")
				fmt.Scan(&programStud)
				switch programStud {
				case 1:
					B[i].program = "S1"
					B[i].jurusan = "Informatika"
				case 2:
					B[i].program = "S1"
					B[i].jurusan = "Sistem Informasi"
				case 3:
					B[i].program = "D4"
					B[i].jurusan = "Informatika"
				case 4:
					B[i].program = "D4"
					B[i].jurusan = "Sistem Informasi"
				case 5:
					B[i].program = "D3"
					B[i].jurusan = "Informatika"
				case 6:
					B[i].program = "D3"
					B[i].jurusan = "Sistem Informasi"
				}

				fmt.Print("Masukan nilai anda: ")
				fmt.Scan(&B[i].nilaiTes)
				fmt.Println()

				i += 1
				fmt.Println()
			}
		}
		*n = i
	case "edit":
		fmt.Print("Suku kata dari nama yang akan dimasukan (1/2/3): ")
		fmt.Scan(&suku)
		switch suku {
		case 1:
			fmt.Print("Nama mahasiswa yang akan diganti:  ")
			fmt.Scan(&edit1)
			fmt.Printf("Masukan data baru\n")
			for i := 0; i < *n; i++ {
				if A[i].namaMahasiswa1 == edit1 {
					fmt.Printf("Nama			     : ")
					fmt.Scan(&edit1)

					fmt.Printf("Nomor Induk Kependudukan     : ")
					fmt.Scan(&nomorIndukEdit)

					fmt.Print("Tanggal Lahir                : ")
					fmt.Scan(&tanggalLahirEdit, &bulanLahirEdit, &tahunLahirEdit)

					fmt.Print("Umur			     : ")
					fmt.Scan(&umurMahasiswaEdit)

					fmt.Printf("Asal Kota (Ex : Bandung)     : ")
					fmt.Scan(&kotaTinggalEdit)

					fmt.Printf("Alamat Rumah (2 suku kata)   : ")
					fmt.Scan(&namaJalan1Edit, &namajalan2Edit)

					fmt.Println()

					fmt.Println("Masukkan Data Registrasi: ")
					menuJurusan()
					fmt.Print("Pilih jurusan: ")
					fmt.Scan(&programStud)
					switch programStud {
					case 1:
						B[i].program = "S1"
						B[i].jurusan = "Informatika"
					case 2:
						B[i].program = "S1"
						B[i].jurusan = "Sistem Informasi"
					case 3:
						B[i].program = "D4"
						B[i].jurusan = "Informatika"
					case 4:
						B[i].program = "D4"
						B[i].jurusan = "Sistem Informasi"
					case 5:
						B[i].program = "D3"
						B[i].jurusan = "Informatika"
					case 6:
						B[i].program = "D3"
						B[i].jurusan = "Sistem Informasi"
					}

					fmt.Print("Masukan nilai anda: ")
					fmt.Scan(&nilaiTesEdit)
					fmt.Println()

					A[i].namaMahasiswa1 = edit1
					A[i].nomorInduk = nomorIndukEdit
					A[i].kotaTinggal = kotaTinggalEdit
					A[i].tanggalLahir = tahunLahirEdit
					A[i].bulanLahir = bulanLahirEdit
					A[i].tahunLahir = tahunLahirEdit
					A[i].namaJalan1 = namaJalan1Edit
					A[i].namajalan2 = namajalan2Edit
					A[i].umurMahasiswa = umurMahasiswaEdit
					B[i].nilaiTes = nilaiTesEdit
				}
			}
		case 2:
			fmt.Print("Nama mahasiswa yang akan diganti:  ")
			fmt.Scan(&edit1, &edit2)
			fmt.Printf("Masukan data baru\n")
			for i := 0; i < *n; i++ {
				if A[i].namaMahasiswa1 == edit1 && A[i].namaMahasiswa2 == edit2 {
					fmt.Printf("Nama			     : ")
					fmt.Scan(&edit1, &edit2)

					fmt.Printf("Nomor Induk Kependudukan     : ")
					fmt.Scan(&nomorIndukEdit)

					fmt.Print("Tanggal Lahir                : ")
					fmt.Scan(&tanggalLahirEdit, &bulanLahirEdit, &tahunLahirEdit)

					fmt.Print("Umur			     : ")
					fmt.Scan(&kotaTinggalEdit)

					fmt.Printf("Asal Kota (Ex : Bandung)     : ")
					fmt.Scan(&kotaTinggalEdit)

					fmt.Printf("Alamat Rumah (2 suku kata)   : ")
					fmt.Scan(&namaJalan1Edit, &namajalan2Edit)

					fmt.Println()

					fmt.Println("Masukkan Data Registrasi: ")
					menuJurusan()
					fmt.Print("Pilih jurusan: ")
					fmt.Scan(&programStud)
					switch programStud {
					case 1:
						B[i].program = "S1"
						B[i].jurusan = "Informatika"
					case 2:
						B[i].program = "S1"
						B[i].jurusan = "Sistem Informasi"
					case 3:
						B[i].program = "D4"
						B[i].jurusan = "Informatika"
					case 4:
						B[i].program = "D4"
						B[i].jurusan = "Sistem Informasi"
					case 5:
						B[i].program = "D3"
						B[i].jurusan = "Informatika"
					case 6:
						B[i].program = "D3"
						B[i].jurusan = "Sistem Informasi"
					}

					fmt.Print("Masukan nilai anda: ")
					fmt.Scan(&nilaiTesEdit)
					fmt.Println()

					A[i].namaMahasiswa1 = edit1
					A[i].nomorInduk = nomorIndukEdit
					A[i].kotaTinggal = kotaTinggalEdit
					A[i].tanggalLahir = tahunLahirEdit
					A[i].bulanLahir = bulanLahirEdit
					A[i].tahunLahir = tahunLahirEdit
					A[i].namaJalan1 = namaJalan1Edit
					A[i].namajalan2 = namajalan2Edit
					A[i].umurMahasiswa = umurMahasiswaEdit
					B[i].nilaiTes = nilaiTesEdit
				}
			}
		case 3:
			fmt.Print("Nama mahasiswa yang akan diganti:  ")
			fmt.Scan(&edit1, &edit2, &edit3)
			fmt.Printf("Masukan data baru\n")
			for i := 0; i < *n; i++ {
				if A[i].namaMahasiswa1 == edit1 && A[i].namaMahasiswa2 == edit2 && A[i].namaMahasiswa3 == edit3 {
					fmt.Printf("Nama			     : ")
					fmt.Scan(&edit1, &edit2, &edit3)

					fmt.Printf("Nomor Induk Kependudukan     : ")
					fmt.Scan(&nomorIndukEdit)

					fmt.Print("Tanggal Lahir                : ")
					fmt.Scan(&tanggalLahirEdit, &bulanLahirEdit, &tahunLahirEdit)

					fmt.Print("Umur			     : ")
					fmt.Scan(&kotaTinggalEdit)

					fmt.Printf("Asal Kota (Ex : Bandung)     : ")
					fmt.Scan(&kotaTinggalEdit)

					fmt.Printf("Alamat Rumah (2 suku kata)   : ")
					fmt.Scan(&namaJalan1Edit, &namajalan2Edit)

					fmt.Println()

					fmt.Println("Masukkan Data Registrasi: ")
					menuJurusan()
					fmt.Print("Pilih jurusan: ")
					fmt.Scan(&programStud)
					switch programStud {
					case 1:
						B[i].program = "S1"
						B[i].jurusan = "Informatika"
					case 2:
						B[i].program = "S1"
						B[i].jurusan = "Sistem Informasi"
					case 3:
						B[i].program = "D4"
						B[i].jurusan = "Informatika"
					case 4:
						B[i].program = "D4"
						B[i].jurusan = "Sistem Informasi"
					case 5:
						B[i].program = "D3"
						B[i].jurusan = "Informatika"
					case 6:
						B[i].program = "D3"
						B[i].jurusan = "Sistem Informasi"
					}

					fmt.Print("Masukan nilai anda: ")
					fmt.Scan(&nilaiTesEdit)
					fmt.Println()

					A[i].namaMahasiswa1 = edit1
					A[i].nomorInduk = nomorIndukEdit
					A[i].kotaTinggal = kotaTinggalEdit
					A[i].tanggalLahir = tahunLahirEdit
					A[i].bulanLahir = bulanLahirEdit
					A[i].tahunLahir = tahunLahirEdit
					A[i].namaJalan1 = namaJalan1Edit
					A[i].namajalan2 = namajalan2Edit
					A[i].umurMahasiswa = umurMahasiswaEdit
					B[i].nilaiTes = nilaiTesEdit
				}
			}
		}
	case "hapus":
		fmt.Print("Ketikan nomor induk yang akan dihapus: ")
		fmt.Scan(&noIndukHapus)
		hasilFind := findDelete(*A, *B, *n, noIndukHapus) // Memanggil fungsi find delete untuk mencari index data yang akan dihapus //

		if hasilFind != -1 {
			for i := hasilFind; i < *n; i++ {
				A[i] = A[i+1]
				B[i] = B[i+1]
			}
			fmt.Println("Penghapusan berhasil !")
			*n = *n - 1
			fmt.Println("Data hasil penghapusan: ")
			isOutput(*A, *B, *n)
		} else {
			fmt.Println("Penghapusan gagal karena tidak ditemukan data yang cocok")
		}
	}
	var clear string
	fmt.Print("Bersihkan layar (yes/no): ")
	fmt.Scan(&clear)
	if clear == "yes" {
		clearScreen()
	}
}

func isOutput(A tabArr, B tabArr2, n int) {
	/* Memerikan hasil data Mahasiswa yang sudah di daftarkan */
	for i := 0; i < n; i++ {
		fmt.Println("==================================================")
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s %s %s		\n", A[i].namaMahasiswa1, A[i].namaMahasiswa2, A[i].namaMahasiswa3)
		fmt.Printf(" Nomor Induk	: %s 			\n", A[i].nomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", A[i].tanggalLahir, A[i].bulanLahir, A[i].tahunLahir)
		fmt.Printf(" Umur		: %d 				\n", A[i].umurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", A[i].kotaTinggal)
		fmt.Printf(" Alamat		: %s %s 		\n", A[i].namaJalan1, A[i].namajalan2)
		fmt.Printf(" Program Studi	: %s %s\n", B[i].program, B[i].jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", B[i].nilaiTes)
		fmt.Println("==================================================")
		fmt.Println()
	}
	var clear string
	fmt.Print("Bersihkan layar (yes/no): ")
	fmt.Scan(&clear)
	if clear == "yes" {
		clearScreen()
	}
}

func jenisOutput(A tabArr, B tabArr2, n int) {
	var jenis string
	fmt.Print("Pilih Daftar (Nama/NIK/TTL/AsKot/Alamat/Prodi/Nilai): ")
	fmt.Scan(&jenis)
	switch jenis {
	case "Nama":
		for i := 0; i < n; i++ {
			fmt.Println(A[i].namaMahasiswa1, A[i].namaMahasiswa2, A[i].namaMahasiswa3)
		}
	case "NIK":
		for i := 0; i < n; i++ {
			fmt.Println(A[i].nomorInduk)
		}
	case "TTL":
		for i := 0; i < n; i++ {
			fmt.Println(A[i].tanggalLahir, A[i].bulanLahir, A[i].tahunLahir)
		}
	case "AsKot":
		for i := 0; i < n; i++ {
			fmt.Println(A[i].kotaTinggal)
		}
	case "Alamat":
		for i := 0; i < n; i++ {
			fmt.Println(A[i].namaJalan1, A[i].namajalan2)
		}
	case "Prodi":
		for i := 0; i < n; i++ {
			fmt.Println(B[i].program, B[i].jurusan)
		}
	case "Nilai":
		for i := 0; i < n; i++ {
			fmt.Println(B[i].nilaiTes)
		}
	}
	var clear string
	fmt.Print("Bersihkan layar (yes/no): ")
	fmt.Scan(&clear)
	if clear == "yes" {
		clearScreen()
	}
}

func oldestYoungestAge(A tabArr, n int) {
	old := 0
	young := 0
	for i := 0; i < n; i++ {
		if A[i].umurMahasiswa > A[old].umurMahasiswa {
			old = i
		}
		if A[i].umurMahasiswa < A[young].umurMahasiswa {
			young = i
		}
	}

	fmt.Printf("Mahasiswa dengan umur tertua bernama %s %s %s yang berumur %d\n", A[old].namaMahasiswa1, A[old].namaMahasiswa2, A[old].namaMahasiswa3, A[old].umurMahasiswa)
	fmt.Printf("Mahasiswa dengan umur tertua bernama %s %s %s yang berumur %d", A[young].namaMahasiswa1, A[young].namaMahasiswa2, A[young].namaMahasiswa3, A[young].umurMahasiswa)
	fmt.Println()
	var clear string
	fmt.Print("Bersihkan layar (yes/no): ")
	fmt.Scan(&clear)
	if clear == "yes" {
		clearScreen()
	}
}

func findDelete(A tabArr, B tabArr2, n int, x string) int {
	// { Fungsi mengembalikan index data yang akan dihapus } //
	idx := -1
	i := 0
	for i < n && idx == -1 {
		if A[i].nomorInduk == x {
			idx = i
		}
		i += 1
	}
	return idx
}

/* TEST CASE




3
Shadra Mehdi mahdi
1203123123
9 8 2004
20
Bandung
Jl.Saluyu Buruk

3
Muhammad Malik hidayot
12313131
8 7 2005
21
Jakarta
Jl.Bintaro Raya
0

*/
