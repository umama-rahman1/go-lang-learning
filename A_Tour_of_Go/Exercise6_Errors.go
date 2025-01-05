package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt computes the square root of x using Newton's method.
// It returns an error if x is negative.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2 // Starting guess
	prevZ := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prevZ) < 1e-10 {
			break
		}
		prevZ = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
