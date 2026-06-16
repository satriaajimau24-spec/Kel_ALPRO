package logical

import (
	"fmt"
)

func HitungKeuntungan() {
	fmt.Println("\n=== HITUNG KEUNTUNGAN ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	totalDana := 0.0
	totalNilai := 0.0

	for _, inv := range DataInvestasi {
		keuntungan := inv.NilaiKini - inv.Dana
		persen := 0.0
		if inv.Dana > 0 {
			persen = (keuntungan / inv.Dana) * 100
		}

		fmt.Printf("%s: Keuntungan = %.2f (%.2f%%)\n",
			inv.Nama, keuntungan, persen)

		totalDana += inv.Dana
		totalNilai += inv.NilaiKini
	}

	totalKeuntungan := totalNilai - totalDana
	fmt.Printf("\nTotal Investasi: %.2f\n", totalDana)
	fmt.Printf("Total Nilai Kini: %.2f\n", totalNilai)
	fmt.Printf("Total Keuntungan: %.2f\n", totalKeuntungan)
}
