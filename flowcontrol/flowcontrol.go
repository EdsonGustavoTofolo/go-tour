package flowcontrol

import (
	"fmt"
	"math"
	"runtime"
)

func For() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func ForLikeWhile() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func ForForever() {
	sum := 1
	for {
		sum += sum
		if sum > 10 {
			break
		}
	}
	fmt.Println(sum)
}

func ForSlice() {

}

func IfShortDeclaration() {
	if v := math.Pow(3, 2); v < 10 {
		fmt.Println("short declaration")
	}
}

func SwitchExpression() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func DeferDeclaration() {
	defer fmt.Println("Tofolo")

	fmt.Println("Edson")
}
