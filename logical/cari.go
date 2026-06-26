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

	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")

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

// ===============================
// SEQUENTIAL SEARCH
// ===============================

func sequentialSearch() {

	kataKunci := strings.ToLower(GetInput("Masukkan kata kunci: "))

	ditemukan := false

	for _, inv := range DataInvestasi {

		if strings.Contains(strings.ToLower(inv.Nama), kataKunci) ||
			strings.Contains(strings.ToLower(inv.Jenis), kataKunci) {

			fmt.Printf("\nID: %s\n", inv.ID[:8])
			fmt.Printf("Nama: %s\n", inv.Nama)
			fmt.Printf("Jenis: %s\n", inv.Jenis)
			fmt.Printf("Dana: %.2f\n", inv.Dana)
			fmt.Printf("Nilai Kini: %.2f\n", inv.NilaiKini)

			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Investasi tidak ditemukan.")
	}
}

// ===============================
// BINARY SEARCH
// ===============================

func binarySearch() {

	namaCari := strings.ToLower(GetInput("Masukkan nama aset: "))

	// Binary Search membutuhkan data yang sudah terurut
	sort.Slice(DataInvestasi, func(i, j int) bool {
		return strings.ToLower(DataInvestasi[i].Nama) < strings.ToLower(DataInvestasi[j].Nama)
	})

	kiri := 0
	kanan := len(DataInvestasi) - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		namaTengah := strings.ToLower(DataInvestasi[tengah].Nama)

		if namaTengah == namaCari {

			inv := DataInvestasi[tengah]

			fmt.Println("\nData ditemukan!")
			fmt.Printf("ID: %s\n", inv.ID[:8])
			fmt.Printf("Nama: %s\n", inv.Nama)
			fmt.Printf("Jenis: %s\n", inv.Jenis)
			fmt.Printf("Dana: %.2f\n", inv.Dana)
			fmt.Printf("Nilai Kini: %.2f\n", inv.NilaiKini)

			return

		} else if namaCari < namaTengah {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1

		}
	}

	fmt.Println("Investasi tidak ditemukan.")
}