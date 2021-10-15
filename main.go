package main

import (
	"fmt"
	"math"

	myRect "github.com/chillrend/myRectCalcModule"
)

func main() {

	circle := struct {
		dia, rad float64
	}{
		dia: 10,
		rad: 5,
	}

	circleArea := func() {
		fmt.Println("Luas Lingkaran: ", math.Pi*(circle.rad*circle.rad))
	}

	circleArea()

	p := myRect.Rect{Len: 5, Wid: 5}

	area, len, wid := p.CalAndReturn()

	fmt.Println("Panjang Kotak: ", len)
	fmt.Println("Lebar Kotak: ", wid)
	fmt.Println("Luas Kotak: ", area)
}
