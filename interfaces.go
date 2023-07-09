package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// For rectangle
type rectangle struct {
	length float64
	breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func (r rectangle) perimeter() float64 {
	return 2 *(r.length * r.breadth)
}

func main() {
	fmt.Println("\nShape: circle")
	c1 := circle{
		radius: 3,
	}

	checkInstance(c1)


	// For rectangle
	fmt.Println("\nShape: rectangle")
	r1 := rectangle{
		length: 3,
		breadth: 5,
	}
	
	checkInstance(r1)

}



// Checking the instance of Shape
func checkInstance(s Shape){
	switch type_of_instance := s.(type){
	case circle:
		fmt.Println(type_of_instance.area())
		break
	case rectangle:
		fmt.Println(type_of_instance.area())
	default:
		fmt.Println("Instance unknown")
	}
	
}
