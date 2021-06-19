package main

import "fmt"
import "math"

type geo interface {
	area() float64
	extent() float64
	volume() float64
}

type rectangle struct {
	width, height float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type cuboid struct {
	width, height, length float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) extent() float64 {
	return 2*r.width + 2*r.height
}

func (r rectangle) volume() float64 {
	return 0
}

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) extent() float64 {
	return 4 * s.length
}

func (s square) volume() float64 {
	return 0
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) extent() float64 {
	return math.Pi * (c.radius + c.radius)
}

func (c circle) volume() float64 {
	return 0
}

func (c cuboid) area() float64 {
	return 2 * (c.width*c.height + c.width*c.length + c.length*c.height)
}
func (c cuboid) extent() float64 {
	return c.width*4 + c.height*4 + c.length*4
}
func (c cuboid) volume() float64 {
	return c.width * c.height * c.length
}

func geocalc(g geo) {
	fmt.Printf("%#v\t%#v\t%#v\t%#v\n", g, g.area(), g.extent(), g.volume())
}

func main() {
	r := rectangle{width: 2, height: 3}
	s := square{length: 3}
	c := circle{radius: 4}
	q := cuboid{width: 3, height: 2, length: 4}

	fmt.Printf("a rectangle with a width of %vm and a height of %vm has a area of %vm² and a extent of %vm\n", r.width, r.height, r.area(), r.extent())
	fmt.Printf("a square with a side length of %vm has a area of %vm² and a extent of %vm\n", s.length, s.area(), s.extent())
	fmt.Printf("a circle with a radius of %vm has a area of %vm² and a extent of %vm\n", c.radius, c.area(), c.extent())
	fmt.Printf("a rectangular cuboid with a width of %vm, a height of %vm and a length of %vm has a area of %vm² and a volume of %vm³\n", q.width, q.height, q.length, q.area(), q.volume())

	geocalc(r)
	geocalc(q)
}
