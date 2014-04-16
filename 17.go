package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println(needInt(Big)) // ./17.go:16: constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Big))
}
