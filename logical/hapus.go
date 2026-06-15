package logical

import (
	"fmt"
	"strings"
)

func HapusInvestasi() {
	fmt.Println("\n=== HAPUS INVESTASI ===")

	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data investasi")
		return
	}

	TampilkanInvestasi()

	idx := CariIndexInvestasi(GetInput("Masukkan ID investasi yang akan dihapus: "))
	if idx == -1 {
		fmt.Println("Investasi tidak ditemukan")
		return
	}

	fmt.Printf("Yakin ingin menghapus '%s'? (y/n): ", DataInvestasi[idx].Nama)
	if konfirmasi := strings.ToLower(GetInput("")); konfirmasi != "y" && konfirmasi != "ya" {
		fmt.Println("Penghapusan dibatalkan")
		return
	}

	DataInvestasi = append(DataInvestasi[:idx], DataInvestasi[idx+1:]...)
	SimpanData()
	fmt.Println("✓ Investasi berhasil dihapus dan disimpan!")
}