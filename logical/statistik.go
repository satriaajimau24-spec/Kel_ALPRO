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

	saham, obligasi, reksadana := 0, 0, 0
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

	type Item struct {
		teks  string
		nilai int
	}

	list := []Item{
		{"Saham: %d\n", saham},
		{"Obligasi: %d\n", obligasi},
		{"ReksaDana: %d\n", reksadana},
	}

	// Logika Insertion Sort
	for i := 1; i < 3; i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j].nilai < key.nilai {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}

	for _, item := range list {
		fmt.Printf(item.teks, item.nilai)
	}
}
