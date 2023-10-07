package main

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