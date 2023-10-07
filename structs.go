package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int	
	message string
}

type car struct {
	Make string
	Model string
	Height int
	Weight int

	// Wheel is a field containd an anonymous struct
	Wheel struct {
		Radius int
		Spokes int
	}
}

type rect struct {
	width int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

type truck struct {
	// Embed the car struct, which means that the fields of car are now
	// fields of truck, and truck can access them directly without having
	// to use the dot notation to access them like truck.car.Make but can 
	// instead just use truck.Make (this is not inheritance, 
	// go does not support inheritance)
	car
	bedSize int
}

func loadTruck(){
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			Make: "Ford",
			Model: "F150",
			Height: 5,
			Weight: 10,
			Wheel: struct {
				Radius int
				Spokes int
			}{Radius: 5, Spokes: 10},
		},
	}

	fmt.Println(lanesTruck)
	fmt.Println(lanesTruck.Make)
	fmt.Println(lanesTruck.Model)
	fmt.Println(lanesTruck.Height)
	fmt.Println(lanesTruck.Weight)
	fmt.Println(lanesTruck.Wheel)
}