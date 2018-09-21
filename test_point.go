package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	a1 := Point{
		X: 1,
		Y: 2,
	}
	a2 := Point{
		X: 3,
		Y: 4,
	}
	a3 := Point{
		X: 5,
		Y: 6,
	}
	p1 := Path{a1, a2}
	fmt.Println(p1)
	p1.TranslateBy(a3, true)
	fmt.Println(p1)
}
