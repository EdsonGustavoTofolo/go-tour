package basics

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func GetConstant() float64 {
	return Pi
}
