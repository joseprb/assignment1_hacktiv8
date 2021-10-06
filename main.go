package main

// import library yang dibutuhkan
import (
	"fmt"
	"os"
	"strconv"
)

// deklarasi struct peserta
type peserta struct {
	nama, alamat      string
	pekerjaan, alasan string
}

// deklarasi array list peserta dengan panjang 10
var list_peserta [10]peserta

// function untuk mencetak identitas peserta
// sesuai dengan nomor absen yang dimasukkan
func print_peserta(idx int) {
	fmt.Printf("Nama\t\t: %v\n", list_peserta[idx-1].nama)
	fmt.Printf("Alamat\t\t: %v\n", list_peserta[idx-1].alamat)
	fmt.Printf("Pekerjaan\t: %v\n", list_peserta[idx-1].pekerjaan)
	fmt.Printf("Alasan\t\t: %v\n", list_peserta[idx-1].alasan)
}

// function untuk mengisi array dengan identitas 10 peserta
func insert_peserta() {
	list_peserta[0].nama = "FIRMAN AULIA INSANI"
	list_peserta[0].alamat = "Jakarta"
	list_peserta[0].pekerjaan = "Mahasiswa"
	list_peserta[0].alasan = "dapat SKS"

	list_peserta[1].nama = "ANDRI NUR HIDAYATULLOH"
	list_peserta[1].alamat = "Bandung"
	list_peserta[1].pekerjaan = "Mahasiswa"
	list_peserta[1].alasan = "dapat uang saku"

	list_peserta[2].nama = "I KOMANG WIDNYANA"
	list_peserta[2].alamat = "Bali"
	list_peserta[2].pekerjaan = "Mahasiswa"
	list_peserta[2].alasan = "dapat uang ilmu"

	list_peserta[3].nama = "ERICO"
	list_peserta[3].alamat = "Semarang"
	list_peserta[3].pekerjaan = "Mahasiswa"
	list_peserta[3].alasan = "dapat uang saku dan SKS"

	list_peserta[4].nama = "JOSE YOLANDA PURBA"
	list_peserta[4].alamat = "Medan"
	list_peserta[4].pekerjaan = "Mahasiswa"
	list_peserta[4].alasan = "dapat uang saku dan ilmu"

	list_peserta[5].nama = "ANDRI KUWITO"
	list_peserta[5].alamat = "Surabaya"
	list_peserta[5].pekerjaan = "Mahasiswa"
	list_peserta[5].alasan = "dapat sks dan ilmu"

	list_peserta[6].nama = "SANDY BUDIMAN"
	list_peserta[6].alamat = "Magelang"
	list_peserta[6].pekerjaan = "Mahasiswa"
	list_peserta[6].alasan = "pengalaman"

	list_peserta[7].nama = "RAFLI ANDREANSYAH"
	list_peserta[7].alamat = "Solo"
	list_peserta[7].pekerjaan = "Mahasiswa"
	list_peserta[7].alasan = "mengisi waktu luang"

	list_peserta[8].nama = "TAUFIQ HIDAYAH"
	list_peserta[8].alamat = "Bekasi"
	list_peserta[8].pekerjaan = "Mahasiswa"
	list_peserta[8].alasan = "mau belajar golang"

	list_peserta[9].nama = "MELVITA SARI"
	list_peserta[9].alamat = "Banten"
	list_peserta[9].pekerjaan = "Mahasiswa"
	list_peserta[9].alasan = "mau belajar web service"
}

func main() {
	// panggil fungsi untuk mengisi array list peserta
	insert_peserta()

	// jika argumen ada
	if len(os.Args) > 1 {
		args, _ := strconv.Atoi(os.Args[1]) // convert argumen ke int
		if args > 0 && args < 11 {          // jika nomor absen valid
			print_peserta(args)
		} else {
			fmt.Println("Nomor absen salah")
		}
	} else {
		fmt.Println("Masukkan nomor absen peserta dalam argument")
	}
}
