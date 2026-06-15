package logical

import (
	"fmt"
	"sort"
)

func UrutkanInvestasi() {
	fmt.Println("\n=== URUTKAN INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Jenis")
	fmt.Println("3. Berdasarkan Dana")
	fmt.Println("4. Berdasarkan Nilai")

	pilihan := GetInput("Pilih pengurutan (1-4): ")

	switch pilihan {
	case "1":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].Nama < DataInvestasi[j].Nama
		})
	case "2":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].Jenis < DataInvestasi[j].Jenis
		})
	case "3":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].Dana < DataInvestasi[j].Dana
		})
	case "4":
		sort.Slice(DataInvestasi, func(i, j int) bool {
			return DataInvestasi[i].NilaiKini < DataInvestasi[j].NilaiKini
		})
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Println("✓ Data berhasil diurutkan")
	TampilkanInvestasi()
}