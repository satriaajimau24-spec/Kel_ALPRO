package logical

import (
	"fmt"
)

func Statistik() {
	fmt.Println("\n=== STATISTIK ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	fmt.Printf("Total Investasi: %d\n", len(DataInvestasi))

	saham := 0
	obligasi := 0
	reksadana := 0

	for _, inv := range DataInvestasi {
		switch inv.Jenis {
		case "Saham":
			saham++
		case "Obligasi":
			obligasi++
		case "ReksaDana":
			reksadana++
		}
	}

	fmt.Printf("Saham: %d\n", saham)
	fmt.Printf("Obligasi: %d\n", obligasi)
	fmt.Printf("ReksaDana: %d\n", reksadana)
}