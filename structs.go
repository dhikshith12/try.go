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

// interface{} is an empty interface, which means that it can be used to
// represent any type of data, in fact every type implements the empty
// interface, so any type can be passed into a function that takes an

type expense interface {
	cost() float64
}

type printer interface {
	print() 
}

type email struct {
	isSubscribed bool
	body string
}

type sms struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}

func (e email) cost() float64 {
	x := float64(len(e.body))
	if !e.isSubscribed {
		return 0.05 * x;
	}
	return 0.01 * x;
}

func (s sms) cost() float64 {
	x := float64(len(s.body))
	if !s.isSubscribed {
		return 0.05 * x;
	}
	return 0.01 * x;
}

func (e email) print() {
	fmt.Println(e.body)
}

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int, err error)
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.body, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.body, s.cost()
	}
	return "", 0
}

func printNumericalValue(num interface{}) {
	switch v:= num.(type){
	case int:
		fmt.Println(v)
	case float64:
		fmt.Println(v)
	default:
		fmt.Println("unknown type")
	}
}