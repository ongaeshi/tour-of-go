// Adavanced Exercise:
// Complex cube roots
package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)

	for {
		pz := z

		// http://tour.golang.org/#50
		// z = z - (z * z * z - x) / 3 * (z * z)
		
		// Here is correct?
		z = (2 * z + x / (z * z)) / 3.0

		if cmplx.Abs(pz - z) < 1e-10 {
			break
		}
	}

	return z
}

func main() {
	cubeRoot := Cbrt(2)
	fmt.Println(cubeRoot, cmplx.Pow(cubeRoot, 3))
}
