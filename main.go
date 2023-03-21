package main

import (
	// "sync"
	"fmt"
	"strings"
	"time"
)

// fmt package is for: print messages, collect user input, write into file
var conferenceName = "Bernards Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50

// array has to define the size and type ===>  var bookings [50]string
// slice is an abstraction of an Array, no need to define the size, better option
// var bookings []map[string]string

var bookings = make([]UserData, 0)

// structs is like lightweight classes
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// uint8 => whole number from 0 to 255, aka: unsigned 8-bit integers
// int8 => whole number from -128 to 127, aka: signed 8-bit integers

// Waitgroup waits for the launched goroutine to finish
// var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	// %T is reference for the type
	// %v is reference for the value
	// fmt.Printf("conferenceName is: %T, remainingTickets is: %T\n", conferenceName, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTickets(userTickets, firstName, lastName, email)

			// "go..." - starts a new goroutine, is a lightweight thread managed by the Go runtime(runs in the background)
			// Add: sets the number of goroutines to wait for
			// wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// 'range'  ===> iterates over elements
			// for arrays and slices , range provides the index and value
			// if no need the index, can use a "Blank identifier" - to ignore a variable ===>   _

			firstNames := getFirstNames()
			fmt.Printf("These are the first names of the booking: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next time!!!")
				// break will end the loop
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address isinvalid")
			}
			if !isValidTicketNumber {
				fmt.Println("numbers of tickets you entered is invalid")
			}
			// fmt.Printf("We onle have %v tickets left! So you can't book %v tickets. \n", remainingTickets, userTickets)
			// continue will only break the current loop
		}
		// Wait: Blocks until the WaitGroup counter is 0
		// wg.Wait()
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!!!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// string.Fields() ===> white space sepatrator, like split in JS
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"]) ====> if it is a map

		// if it is a struct
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// in go functions can return multiple values
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	// Scan(userName)  ===> save users's value in "userName" variable
	// & is a pointer ===> reference the memory address

	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets
	// fill up an array ==>    bookings[0] = firstName + " " + lastName
	// append ==> built in function to adding a value for a slice

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// creating an empty map with "make"
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank You %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	// Done: Decrements the WaitGroup counter by 1
	// wg.Done()
}
