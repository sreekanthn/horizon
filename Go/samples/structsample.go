package main

import "fmt"
import "math"

type Point struct {
	X, Y float64
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) Scaleby(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{10, 25}
	q := Point{5, 22}
	var r *Point
	dist := Distance

	fmt.Println(p.Distance(q))
	fmt.Println(Distance(p, q))
	fmt.Println(dist(p, q))
	r = &p
	r.Scaleby(5)

	scale := (*Point).Scaleby
	scale(&q, 4)
	fmt.Println(p.Distance(q))
}
