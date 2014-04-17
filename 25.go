// Exercise: Loops and Functions
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; true; i++ {
		pz := z
		z = z - (z * z - x) / 2 * z

		if math.Abs(pz - z) < 0.0001 {
			fmt.Println(i, ": ", z)
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
