package logical

type Investasi struct {
	ID        string
	Nama      string
	Jenis     string
	Dana      float64
	NilaiKini float64
}

var DataInvestasi []Investasi

const DataFile = "logical/dataInvestasi.txt"