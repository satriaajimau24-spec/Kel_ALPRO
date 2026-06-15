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

	//Selection Sort (terbesar -> terkecil)
	type RingkasanAset struct {
		Label string
		Nilai float64
	}

	listLaporan := []RingkasanAset{
		{Label: "Saham    ", Nilai: saham},
		{Label: "Obligasi ", Nilai: obligasi},
		{Label: "ReksaDana", Nilai: reksadana},
	}

	n := len(listLaporan)
	for i := 0; i < n-1; i++ {
		idxMaksimal := i
		for j := i + 1; j < n; j++ {
			if listLaporan[j].Nilai > listLaporan[idxMaksimal].Nilai {
				idxMaksimal = j
			}
		}
		listLaporan[i], listLaporan[idxMaksimal] = listLaporan[idxMaksimal], listLaporan[i]
	}

	for _, laporan := range listLaporan {
		fmt.Printf("%s : %.2f (%.1f%%)\n", laporan.Label, laporan.Nilai, persen(laporan.Nilai))
	}
}
