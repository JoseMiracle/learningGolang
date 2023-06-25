package main

import (
	"fmt"
)

type car struct {
	name   string
	year   int
	color  string
	isFast bool
}

func (v vehicle) Details() string {
	return fmt.Sprintf("The name of the vehicle is %s, year of production is %d, its color is %s, and isFast is %t",
		v.car.name,
		v.car.year,
		v.car.color,
		v.car.isFast,
	)
}

type vehicle struct {
	isLong bool
	car
}

func main() {
	// Struct Example 1
	car1 := car{
		name:   "BMW",
		year:   2021,
		color:  "red",
		isFast: true,
	}

	fmt.Println(car1.name)

	// Struct Example 2
	vehicle1 := vehicle{
		isLong: true,
		car: car{
			name:   "limousine",
			year:   2023,
			color:  "yellow",
			isFast: true,
		},
	}

	fmt.Println(vehicle1.year)
	fmt.Println(vehicle1.Details())
}
