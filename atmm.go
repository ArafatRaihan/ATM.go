package main

import (
	"fmt"
)

// Struktur data untuk akun mahasiswa
type Akun struct {
	Username string
	NPM      string
	Saldo    int64
	Histori  []string
}

// Daftar akun mahasiswa
var akunMahasiswa = map[string]*Akun{
	"arafat": {
		Username: "arafat",
		NPM:      "17022006",
		Saldo:    3500000,
		Histori:  []string{},
	},
}

func main() {
	var username, password string

	fmt.Println("Masukkan Username dan Password:")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	// Memverifikasi username dan password
	if akun, ok := akunMahasiswa[username]; ok && akun.NPM == password {
		for {
			fmt.Println("\nPilih Menu:")
			fmt.Println("1. Lihat Informasi Akun")
			fmt.Println("2. Lihat Saldo")
			fmt.Println("3. Tambah Saldo")
			fmt.Println("4. Tarik Saldo")
			fmt.Println("5. Histori Transaksi")
			fmt.Println("6. Keluar dari Program")
			fmt.Println("7. Ubah Password")

			var pilihan int
			fmt.Print("Pilih Menu: ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case 1:
				fmt.Printf("\nInformasi Akun:\nUsername: %s\nNPM: %s\n", akun.Username, akun.NPM)

			case 2:
				fmt.Printf("\nSaldo Anda: Rp %d\n", akun.Saldo)

			case 3:
				var tambahSaldo int64
				fmt.Print("Masukkan jumlah saldo yang ingin ditambahkan: ")
				fmt.Scanln(&tambahSaldo)
				if tambahSaldo < 0 {
					fmt.Println("Jumlah yang ingin ditambahkan tidak bisa negatif.")
					break
				}
				akun.Saldo += tambahSaldo
				akun.Histori = append(akun.Histori, fmt.Sprintf("Saldo ditambahkan sebesar Rp %d", tambahSaldo))
				fmt.Printf("\nSaldo Anda telah ditambahkan sebesar Rp %d. Saldo baru: Rp %d\n", tambahSaldo, akun.Saldo)

			case 4:
				var tarikSaldo int64
				fmt.Print("Masukkan jumlah saldo yang ingin ditarik: ")
				fmt.Scanln(&tarikSaldo)
				if tarikSaldo < 0 {
					fmt.Println("Jumlah yang ingin ditarik tidak bisa negatif.")
					break
				}
				if akun.Saldo-tarikSaldo >= 0 {
					akun.Saldo -= tarikSaldo
					akun.Histori = append(akun.Histori, fmt.Sprintf("Saldo ditarik sebesar Rp %d", tarikSaldo))
					fmt.Printf("\nSaldo Anda telah ditarik sebesar Rp %d. Saldo baru: Rp %d\n", tarikSaldo, akun.Saldo)
				} else {
					fmt.Println("Saldo tidak cukup untuk melakukan penarikan.")
				}
			case 5:
				fmt.Println("\nHistori Transaksi:")
				if len(akun.Histori) == 0 {
					fmt.Println("Tidak ada transaksi.")
				} else {
					for _, transaksi := range akun.Histori {
						fmt.Println(transaksi)
					}
				}
			case 6:
				fmt.Println("\nTerima kasih telah menggunakan ATM!")
				return

			case 7:
				fmt.Print("Masukkan password baru: ")
				var newPassword string
				fmt.Scanln(&newPassword)
				akun.NPM = newPassword
				fmt.Println("Password berhasil diubah!")
			default:
				fmt.Println("Pilihan tidak tersedia. Silakan pilih menu yang tersedia.")
			}
		}
	} else {
		fmt.Println("Username atau Password salah. Silakan coba lagi.")
	}
}
