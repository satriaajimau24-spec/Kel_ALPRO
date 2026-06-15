// logical/laporan.go
package logical

import (
	"fmt"
)

func LaporanPortofolio() {
	fmt.Println("\n=== LAPORAN PORTOFOLIO ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	saham := 0.0
	obligasi := 0.0
	reksadana := 0.0

	for _, inv := range DataInvestasi {
		switch inv.Jenis {
		case "Saham":
			saham += inv.NilaiKini
		case "Obligasi":
			obligasi += inv.NilaiKini
		case "ReksaDana":
			reksadana += inv.NilaiKini
		}
	}

	total := saham + obligasi + reksadana

	persen := func(nilai float64) float64 {
		if total == 0 {
			return 0
		}
		return (nilai / total) * 100
	}

	fmt.Printf("\nTotal Portofolio: %.2f\n", total)
	fmt.Printf("Saham    : %.2f (%.1f%%)\n", saham, persen(saham))
	fmt.Printf("Obligasi : %.2f (%.1f%%)\n", obligasi, persen(obligasi))
	fmt.Printf("ReksaDana: %.2f (%.1f%%)\n", reksadana, persen(reksadana))
}