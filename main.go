package main

import (
	"fmt"
	"tubes/logical"
)

func main() {
	logical.MuatData()

	for {
		logical.ClearScreen()
		fmt.Println("=== APLIKASI MANAJEMEN INVESTASI ===")
		fmt.Println("1. Tambah Investasi")
		fmt.Println("2. Ubah Investasi")
		fmt.Println("3. Hapus Investasi")
		fmt.Println("4. Hitung Keuntungan")
		fmt.Println("5. Cari Investasi")
		fmt.Println("6. Urutkan Investasi")
		fmt.Println("7. Laporan Portofolio")
		fmt.Println("8. Statistik")
		fmt.Println("9. Keluar")

		switch logical.GetInput("\nPilih menu (1-9): ") {
		case "1":
			logical.TambahInvestasi()
		case "2":
			logical.UbahInvestasi()
		case "3":
			logical.HapusInvestasi()
		case "4":
			logical.HitungKeuntungan()
		case "5":
			logical.CariInvestasi()
		case "6":
			logical.UrutkanInvestasi()
		case "7":
			logical.LaporanPortofolio()
		case "8":
			logical.Statistik()
		case "9":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln()
	}
}