package logical

import (
	"fmt"
	"sort"
	"strings"
)

func CariInvestasi() {
	fmt.Println("\n=== CARI INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	fmt.Println("1. Sequential Search (Nama/Jenis)")
	fmt.Println("2. Binary Search (ID)")

	pilihan := GetInput("Pilih metode pencarian (1-2): ")

	switch pilihan {
	case "1":
		sequentialSearch()
	case "2":
		binarySearch()
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

// =====================================
// SEQUENTIAL SEARCH
// =====================================

func sequentialSearch() {

	kataKunci := strings.ToLower(GetInput("Masukkan nama atau jenis aset: "))

	ditemukan := false

	for _, inv := range DataInvestasi {

		if strings.Contains(strings.ToLower(inv.Nama), kataKunci) ||
			strings.Contains(strings.ToLower(inv.Jenis), kataKunci) {

			tampilkanData(inv)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Investasi tidak ditemukan.")
	}
}

// =====================================
// BINARY SEARCH
// =====================================

func binarySearch() {

	// Urutkan berdasarkan ID
	sort.Slice(DataInvestasi, func(i, j int) bool {
		return DataInvestasi[i].ID < DataInvestasi[j].ID
	})

	fmt.Println("\nDaftar ID Investasi:")

	for _, inv := range DataInvestasi {
		fmt.Printf("%s - %s\n", inv.ID[:8], inv.Nama)
	}

	idCari := strings.ToLower(GetInput("\nMasukkan 8 karakter ID: "))

	kiri := 0
	kanan := len(DataInvestasi) - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		idTengah := strings.ToLower(DataInvestasi[tengah].ID[:8])

		if idTengah == idCari {

			fmt.Println("\nData ditemukan!")
			tampilkanData(DataInvestasi[tengah])
			return

		} else if idCari < idTengah {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1

		}
	}

	fmt.Println("Investasi tidak ditemukan.")
}

// =====================================
// MENAMPILKAN DATA
// =====================================

func tampilkanData(inv Investasi) {

	fmt.Println("--------------------------------")
	fmt.Printf("ID         : %s\n", inv.ID)
	fmt.Printf("Nama       : %s\n", inv.Nama)
	fmt.Printf("Jenis      : %s\n", inv.Jenis)
	fmt.Printf("Dana       : Rp %.2f\n", inv.Dana)
	fmt.Printf("Nilai Kini : Rp %.2f\n", inv.NilaiKini)
	fmt.Println("--------------------------------")
}