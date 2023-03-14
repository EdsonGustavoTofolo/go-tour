package generics

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Subtrair[T int | float64](m map[string]T) T {
	var sub T
	for _, v := range m {
		sub -= v
	}
	return sub
}

func RunGenerics() {
	mInt := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
	mFloat := map[string]float64{"A": 1.1, "B": 2.2, "C": 3.3, "D": 4.4}
	mMyNumber := map[string]MyNumber{"A": 1, "B": 2, "C": 3, "D": 4}

	println(Soma(mInt), Soma(mFloat), Soma(mMyNumber))
}
