package struct_tubes

type TabPenerimaan struct {
    NamaMahasiswa, NomorInduk 								string
    KotaTinggal, NamaJalan                        			string
    TanggalLahir, BulanLahir, TahunLahir, UmurMahasiswa     int
}

type TabRegis struct {
    Program, JalurSeleksi, Jurusan string
    NilaiTes                       float64
}

type KonfirmasiAdmin struct {
	Konfirmasi bool
}
const NMAX int = 1000

type TabArr [NMAX]TabPenerimaan
type TabArr2 [NMAX]TabRegis
type TabKonf [NMAX]KonfirmasiAdmin