// Fiboacci closure
package main

import "fmt"

func fibonacci() func() int {
	n  := 0
	f0 := 0
	f1 := 1

	return func () int {
		n += 1

		switch n-1 {
		case 0:
			return f0
		case 1:
			return f1
		default:
			fib := f0 + f1
			f0 = f1
			f1 = fib
			return fib
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
