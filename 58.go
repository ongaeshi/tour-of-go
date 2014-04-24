package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
	if (f < 0.0) {
		return 0, ErrNegativeSqrt(f)
	}
	
	
	z := float64(1)

	for i := 0; true; i++ {
		pz := z
		z = z - (z * z - f) / 2 * z

		if math.Abs(pz - z) < 0.0001 {
			// fmt.Println(i, ": ", z)
			break
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
