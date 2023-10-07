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

type rect struct{
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


type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name 			string
	hourlyRate 		int
	hoursPerYear 	int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyRate * c.hoursPerYear
}

type fullTime struct {
	name 		string
	salary 		int
}

func (f fullTime) getName() string {
	return f.name
}

func (f fullTime) getSalary() int {
	return f.salary
}

func printEmployee(e employee) {
	fmt.Println("name:", e.getName(), "salary:", e.getSalary())
}
