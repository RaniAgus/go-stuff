package main

import (
	"fmt"
)

type SqrtError float64

func (e SqrtError) Error() string {
	return fmt.Sprintf("Sqrt doesn't support complex numbers. Given: %v",
		float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x <= 0 {
		return 0, SqrtError(x)
	}

	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
