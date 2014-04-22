// Adavanced Exercise:
// Complex cube roots
package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)

	for i := 0; true; i++ {
		pz := z

		// failed
		// z = z - (z * z * z - x) / 3 * (z * z)

		// correct
		z = (2 * z + x / (z * z)) / 3.0

		if cmplx.Abs(pz - z) < 0.00000001 {
			// fmt.Println(i, ": ", z)
			break
		}
	}

	return z
}

func main() {
	cubeRoot := Cbrt(2)
	fmt.Println(cubeRoot, cmplx.Pow(cubeRoot, 3))
}
