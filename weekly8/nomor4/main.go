package main

import (
	"fmt"
	"math"
)

// circle
// rectangle

type circle struct {
	radius float64
}

type rectangle struct {
	height float64
	width float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) Area() float64{
	return r.height * r.width
}

type shape interface {
	Area() float64
}

func calc(shapes []shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	shapes := []shape{
		circle{radius: 5},
		rectangle{width: 3, height: 4},
		circle{radius: 2},
	}

	fmt.Println("Total area:", calc(shapes))
}

