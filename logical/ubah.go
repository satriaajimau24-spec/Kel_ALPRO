package logical

import (
	"fmt"
)

func UbahInvestasi() {
	fmt.Println("\n=== UBAH INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	TampilkanInvestasi()

	idx := CariIndexInvestasi(GetInput("Masukkan ID investasi yang akan diubah: "))
	if idx == -1 {
		fmt.Println("Investasi tidak ditemukan")
		return
	}

	fmt.Printf("\nMengubah data: %s\n", DataInvestasi[idx].Nama)
	DataInvestasi[idx].Nama = GetInput("Nama Aset baru: ")
	DataInvestasi[idx].Jenis = GetInput("Jenis Aset baru: ")
	DataInvestasi[idx].Dana = GetFloatInput("Jumlah Dana baru: ")
	DataInvestasi[idx].NilaiKini = GetFloatInput("Nilai Terkini baru: ")

	SimpanData()
	fmt.Println("✓ Investasi berhasil diubah dan disimpan!")
}