package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var seed = int64(time.Now().Nanosecond()) 
	rand.Seed(seed)
	fmt.Println("Seed is", seed, "My favorite number is", rand.Intn(10))
}
