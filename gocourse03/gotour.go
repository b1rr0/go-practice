package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = log(x, y)
		}
		res[y] = row
	}
	return res
}

func log(x, y int) uint8 {
	return uint8(x ^ y)
}

func main() {
	pic.Show(Pic)
}
