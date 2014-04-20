package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    canvas := make([][]uint8, dy)
    
    for y := 0; y < dy; y++ {
        canvas[y] = make([]uint8, dx)
                
        for x := 0; x < dx; x++ {
            canvas[y][x] = uint8(x ^ y)
            // canvas[y][x] = uint8((x + y) / 2)
            // canvas[y][x] = uint8(x * y)
            // canvas[y][x] = uint8(x + y)
        }
    }
    
    return canvas
}

func main() {
    pic.Show(Pic)
}

