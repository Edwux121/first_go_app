package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// stores user information into Map
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wait Group
var wg = sync.WaitGroup{}

func main() {
	//calling greetUsers function
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	//checking user input
	isValidName, isValidEmail, isValidTicketNumber := ValidareUserInput(firstName, lastName, email, userTickets, remainingTickets)

	//checking if we have enough tickets for the user
	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//call function print first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of the bookings: %v\n", firstNames)

		//cheking if there is remaining tickets
		if remainingTickets == 0 {
			//end of the program
			fmt.Println("Our conference is booked out. Come back next time!")
			//break
		}
	} else {
		// !=
		if !isValidName {
			fmt.Println("Your first name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Your email address doesn't contain @ symbol.")
		}
		if !isValidTicketNumber {
			fmt.Println("The number of tickets you've entered is invalid.")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	//taking only firstname from the slice
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	//fmt.Printf("The first names of the bookings: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	//variables for storing user data
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// asking user for the user input
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	//updates remaining tickets with the users bought tickets
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//we update bookings slice with First Name and Last Name
	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}
