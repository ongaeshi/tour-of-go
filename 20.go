// 前後のセミコロンは省略出来る、C言語のwhileに相当
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
