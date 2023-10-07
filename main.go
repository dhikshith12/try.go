package main

import (
	"fmt"
)

/*
	Constants are declared like variables, but with the const keyword.
	Constants can be character, string, boolean, or numeric values.
	Constants cannot be declared using the := syntax.
	
	:= is for variable declaration and initialization.
	it is shorthand for declaring a variable, and then initializing it.
	it's called the short variable declaration operator or in short, short declaration operator.

*/
func sub(x int, y int) int {
	return x - y;
}

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60

	const secondsInHour = secondsInMinute * minutesInHour


	fmt.Println("number of seconds in an hour: ", secondsInHour)

	msg := fmt.Sprintf("subtract %d from %d: %d", 5, 3, sub(5, 3))

	fmt.Println(msg)

	const name = "Saul Goodman"
	const openRate = 0.75

	msg2 := fmt.Sprintf("%s has an open rate of %.2f%%", name, openRate*100)

	fmt.Println(msg2)

/*
	var height int
	// take input from user
	fmt.Println("Enter your height in inches: ")
	fmt.Scan(&height)


	if height > 4 {
		fmt.Println("You are tall")
	} else {
		fmt.Println("You are short")
	}
*/

	messageLen := 10
	maxMessageLen := 20


	if messageLen > maxMessageLen {
		fmt.Println("Your message is too long.")
	} else {
		fmt.Println("Your message is fine.")
	}

	// if statement with initialization
	if length := getLength("hello"); length > 5 {
		fmt.Println("That's a long string.")
	} else {
		fmt.Println("That's a short string.")
	}

	// you can't access length here!
	user, err := createUser("Saul", "Goodman", 50)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("user:", user)

	user2, err := createUser("Saul", "Goodman", -10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("user:", user2)

	fmt.Println(dist(add, 5))

	firstName, lastName := getNames()

	fmt.Println(firstName, lastName)

	firstName2, _ := getNames()

	fmt.Println(firstName2)

	x, y := getCoords();

	fmt.Println(x, y)

	for i := 1; i < 100; i++ {
		fmt.Println(yearUntilEvents(i))
	}

	quotient, err := divide(5, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("quotient:", quotient)

	msg3 := messageToSend {
		phoneNumber: 5551234567,
		message: "Hello, world!",
	}

	fmt.Println(msg3)

	fmt.Println("can send message?:", canSendMessage(msg3))

	myCar := struct {
		make string
		model string
	} {
		make: "Ford",
		model: "F150",
	}

	fmt.Println(myCar)

	r := rect{
		width: 5,
		height: 10,
	}

	fmt.Println("area:", r.area())
	loadTruck();

	auth := authentication{
		username: "user",
		password: "pass",
	}

	fmt.Println(auth.autherizationString())
}

var add func(a, b int) int = func(a int, b int) int {
	return a+b;
}


func getLength(s string) int {
	return len(s)
}

func createUser(firstName, lastName string, age int) (string, error) {
	if age < 0 {
		return "", fmt.Errorf("age %d is not valid", age);
	}
	return fmt.Sprintf("%s %s is %d years old", firstName, lastName, age), nil
}

var dist func(func(int, int) int, int) int = func(f func(a int, b int) int, x int) int {
	return f(x, x) + x;
}

func getNames() (string, string) {
	return "Saul", "Goodman"
}

func getCoords() (x, y int) {
	x++;
	y++;
	return;
}

func yearUntilEvents(age int) (
	yearsUntilDriver int,
	yearsUntilDrinking int,
	yearsUntilRetirement int,
) {
	yearsUntilDriver = max(0, 18 - age)
	yearsUntilDrinking = max(0, 21 - age)
	yearsUntilRetirement = max(0, 65 - age)
	return;
}

func divide(dividend, divisor int) (float64, error) {
	// early return
	if divisor == 0 {
		return 0, fmt.Errorf("can't divide by 0")
	}
	return float64(dividend) / float64(divisor), nil
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.message == "" {
		return false
	}
	if mToSend.phoneNumber == 0 {
		return false
	}
	return true;
}