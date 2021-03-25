package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var picture [][]uint8

	for y := 0; y < dy; y++ {
		var col []uint8
		for x := 0; x < dx; x++ {
			col = append(col, uint8((dx-x)*y+(dy-y)*x+(y+x)/2))
		}
		picture = append(picture, col)
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
