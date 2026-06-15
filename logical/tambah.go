package logical

import (
	"crypto/rand"
	"fmt"
)

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func TambahInvestasi() {
	fmt.Println("\n=== TAMBAH INVESTASI ===")

	inv := Investasi{
		ID:        generateID(),
		Nama:      GetInput("Nama Aset: "),
		Jenis:     GetInput("Jenis Aset (Saham/Obligasi/ReksaDana): "),
		Dana:      GetFloatInput("Jumlah Dana: "),
		NilaiKini: GetFloatInput("Nilai Terkini: "),
	}

	DataInvestasi = append(DataInvestasi, inv)
	SimpanData()
	fmt.Println("✓ Investasi berhasil ditambahkan dan disimpan!")
}