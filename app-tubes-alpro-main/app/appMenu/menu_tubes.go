package menu_tubes

import "fmt"

func MenuAdminUser() {
	fmt.Println("==========Login Sebagai===========")
	fmt.Println("||	1. Admin 		||")
	fmt.Println("||	2. Mahasiswa		||")
	fmt.Println("||	3. Exit			||")
	fmt.Println("==================================")
} // Ditampilkan saat pertama kali program dijalankan

func MenuAdmin() {
	fmt.Println("==================================================")
	fmt.Println("Selamat Datang!")
	fmt.Println("==================================================")
	fmt.Println("||	1. Cek Daftar Mahasiswa			||")
	fmt.Println("||	2. Konfirmasi Kelulusan Mahasiswa	||")
	fmt.Println("||	3. Edit Daftar Mahasiswa		||")
	fmt.Println("||	4. Kembali				||")
	fmt.Println("==================================================")
} // Ditampilkan saat pengguna login sebagai admin

func MenuSortingAdmin() {
	fmt.Println("=============Urutkan berdasarkan=============")
	fmt.Println("||	1. Nilai Tes Rendah - Tinggi	   ||")
	fmt.Println("||	2. Nilai Tes Tinggi - Rendah	   ||")
	fmt.Println("=============Cari berdasarkan================")
	fmt.Println("||	3. Jurusan			   ||")
	fmt.Println("||	4. Mahasiswa Diterima		   ||")
	fmt.Println("||	5. Mahasiswa Ditolak		   ||")
	fmt.Println("||	6. Cari Mahasiswa Dengan Umur	   ||")
	fmt.Println("||	7. Kembali  			   ||")
	fmt.Println("=============================================")
} // Ditampilkan saat pengguna memilih menu Cek Pendaftar di halaman admin

func MenuJurusan() {
	fmt.Println("======= Pilih Program Studi=======")
	fmt.Println("|| 1. S1 Informatika		||")
	fmt.Println("|| 2. S1 Sistem Informasi	||")
	fmt.Println("|| 3. D4 Informatika		||")
	fmt.Println("|| 4. D4 Sistem Informasi 	||")
	fmt.Println("|| 5. D3 Informatika		||")
	fmt.Println("|| 6. D3 Sistem Informasi	||")
	fmt.Println("==================================")
	fmt.Println()
} // Ditampilkan setelah mahasiswa mengisi informasi data diri

func MenuUser() {
	fmt.Println("===================================")
	fmt.Println("Selamat Datang,")
	fmt.Println("Portal Pendfaftaran Calon Mahasiswa")
	fmt.Println("===================================")
	fmt.Println("||1. Daftar			||")
	fmt.Println("||2. Cek Status Pendaftaran	||")
	fmt.Println("||3. Cek Status Kelulusan	||")
	fmt.Println("||4. Kembali			||")
	fmt.Println("==================================")
} // Ditampilkan di halaman Mahasiswa setelah mengisi data diri

func SortingJurusan() {
	fmt.Println("======= Pilih Berdasarkan ========")
	fmt.Println("|| 1. S1 Informatika		||")
	fmt.Println("|| 2. S1 Sistem Informasi	||")
	fmt.Println("|| 3. D4 Informatika		||")
	fmt.Println("|| 4. D4 Sistem Informasi 	||")
	fmt.Println("|| 5. D3 Informatika		||")
	fmt.Println("|| 6. D3 Sistem Informasi	||")
	fmt.Println("==================================")
} // Ditampilkan di halaman admin Cek Pendaftar