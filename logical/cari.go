package logical

import (
	"fmt"
	"strings"
)

func CariInvestasi() {
	fmt.Println("\n=== CARI INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	kataKunci := strings.ToLower(GetInput("Masukkan kata kunci pencarian: "))

	ditemukan := false
	for _, inv := range DataInvestasi {
		if strings.Contains(strings.ToLower(inv.Nama), kataKunci) ||
			strings.Contains(strings.ToLower(inv.Jenis), kataKunci) {
			fmt.Printf("ID: %s | Nama: %s | Jenis: %s | Dana: %.2f | Nilai: %.2f\n",
				inv.ID[:8], inv.Nama, inv.Jenis, inv.Dana, inv.NilaiKini)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Investasi tidak ditemukan")
	}
}
