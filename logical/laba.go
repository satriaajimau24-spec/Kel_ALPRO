package logical

import (
	"fmt"
)

func HitungLaba() {
	fmt.Println("\n=== HITUNG KEUNTUNGAN / KERUGIAN ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	totalDana := 0.0
	totalNilai := 0.0

	for _, inv := range DataInvestasi {
		selisih := inv.NilaiKini - inv.Dana
		persen := 0.0

		if inv.Dana > 0 {
			persen = (selisih / inv.Dana) * 100
		}

		if selisih > 0 {
			fmt.Printf("%s: Keuntungan = %.2f (%.2f%%)\n",
				inv.Nama, selisih, persen)
		} else if selisih < 0 {
			fmt.Printf("%s: Kerugian = %.2f (%.2f%%)\n",
				inv.Nama, -selisih, -persen)
		} else {
			fmt.Printf("%s: Impas (0.00%%)\n", inv.Nama)
		}

		totalDana += inv.Dana
		totalNilai += inv.NilaiKini
	}

	totalSelisih := totalNilai - totalDana

	fmt.Printf("\nTotal Investasi : %.2f\n", totalDana)
	fmt.Printf("Total Nilai Kini: %.2f\n", totalNilai)

	if totalSelisih > 0 {
		fmt.Printf("Total Keuntungan: %.2f\n", totalSelisih)
	} else if totalSelisih < 0 {
		fmt.Printf("Total Kerugian: %.2f\n", -totalSelisih)
	} else {
		fmt.Println("Total Impas: 0.00")
	}
}
