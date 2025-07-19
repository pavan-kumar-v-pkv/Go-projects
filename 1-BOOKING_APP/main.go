package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// package level variables

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // slice to hold user data
//var bookings = make([]map[string]string, 0) // slice with dynamic size

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Printf("conferenceTickets is of type: %T, remainingTickets is of type: %T, conferenceName is of type: %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers()

	// infinite loop to keep asking for user input until tickets run out
	// get user input
	firstName, lastName, email, userTickets := getUserInput()
	// validate user input
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		// call function to book ticket
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		// check if the remaining tickets are less than or equal to 0
		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out. Come back next year!")
			// break // exit the infinite loop
		}
	} else {
		if !isValidName {
			fmt.Println("Your first name or last name entered is too short. They each must be at least 2 characters long.")
		}
		if !isValidEmail {
			fmt.Println("Your email address must contain an '@' and a '.'")
		}
		if !isValidTicketNumber {
			fmt.Printf("You can book a maximum of %v tickets. You tried to book %v tickets.\n", remainingTickets, userTickets)
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{} // slice to hold first names of all bookings
	for _, booking := range bookings {
		// var names = strings.Fields(booking) // split the booking string into fields (first name and last name)
		firstNames = append(firstNames, booking.firstName) // append the first name to the slice
	}
	// fmt.Printf("These are all our bookings: %v\n", bookings)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Print("Enter your name: ")
	fmt.Scan(&firstName)

	// ask user for their last name
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask user for their email
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	// ask user for number of tickets
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // for slice with dynamic size
	fmt.Printf("List of bookings is now: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // simulate a delay for sending the ticket
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
