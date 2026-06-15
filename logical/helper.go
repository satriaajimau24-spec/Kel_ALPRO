package logical

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ─── Input ───────────────────────────────────────────────────────────────────

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetFloatInput(prompt string) float64 {
	input := GetInput(prompt)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid, menggunakan 0")
		return 0
	}
	return value
}

// ─── File ─────────────────────────────────────────────────────────────────────

// MuatData membaca dataInvestasi.txt ke slice DataInvestasi saat program mulai.
func MuatData() {
	file, err := os.Open(DataFile)
	if err != nil {
		return // file belum ada, tidak masalah
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) != 5 {
			continue
		}
		dana, _ := strconv.ParseFloat(parts[3], 64)
		nilai, _ := strconv.ParseFloat(parts[4], 64)
		DataInvestasi = append(DataInvestasi, Investasi{
			ID:        parts[0],
			Nama:      parts[1],
			Jenis:     parts[2],
			Dana:      dana,
			NilaiKini: nilai,
		})
	}
}

// SimpanData menulis seluruh DataInvestasi ke dataInvestasi.txt.
func SimpanData() {
	file, err := os.Create(DataFile)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}
	defer file.Close()

	for _, inv := range DataInvestasi {
		fmt.Fprintf(file, "%s|%s|%s|%.2f|%.2f\n",
			inv.ID, inv.Nama, inv.Jenis, inv.Dana, inv.NilaiKini)
	}
}

// ─── Utilitas ─────────────────────────────────────────────────────────────────

// CariIndexInvestasi mencari index berdasarkan ID (bisa ID penuh atau 8 karakter pertama).
func CariIndexInvestasi(idInput string) int {
	idInput = strings.TrimSpace(strings.ToLower(idInput))
	for i, inv := range DataInvestasi {
		idLengkap := strings.ToLower(inv.ID)
		idPendek := idLengkap
		if len(idLengkap) > 8 {
			idPendek = idLengkap[:8]
		}
		if idLengkap == idInput || idPendek == idInput || strings.Contains(idLengkap, idInput) {
			return i
		}
	}
	return -1
}

// TampilkanInvestasi mencetak semua data investasi ke layar.
func TampilkanInvestasi() {
	if len(DataInvestasi) == 0 {
		fmt.Println("Tidak ada data")
		return
	}
	sep := strings.Repeat("-", 72)
	fmt.Println("\nDaftar Investasi:")
	fmt.Println(sep)
	for i, inv := range DataInvestasi {
		fmt.Printf("%d. ID: %s\n", i+1, inv.ID)
		fmt.Printf("   Nama: %s | Jenis: %s\n", inv.Nama, inv.Jenis)
		fmt.Printf("   Dana: Rp %.2f | Nilai Kini: Rp %.2f\n", inv.Dana, inv.NilaiKini)
		fmt.Println(sep)
	}
}