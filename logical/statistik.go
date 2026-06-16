package logical

import "fmt"

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

	teks := []string{"Saham: %d\n", "Obligasi: %d\n", "ReksaDana: %d\n"}
	nilai := []int{saham, obligasi, reksadana}

	// LOGIKA INSERTION SORT

	for i := 1; i < 3; i++ {
		keyNilai := nilai[i]
		keyTeks := teks[i]

		j := i - 1

		for j >= 0 && nilai[j] > keyNilai {
			nilai[j+1] = nilai[j]
			teks[j+1] = teks[j]
			j--
		}

		nilai[j+1] = keyNilai
		teks[j+1] = keyTeks
	}

	for i := 0; i < 3; i++ {
		fmt.Printf(teks[i], nilai[i])
	}
}
