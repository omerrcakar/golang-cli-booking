package main

// go programs are organized into package
// go standart library, provides different core packages for us to use

// a program can only have 1 main function, because you can only have 1 entypoint
import (
	"booking-app/helper"
	"fmt" // standart Format package
	"strings"
)

// packace level variable / we can use normal variable level syntax
// define variable as 'local' as possible
// create the varible where you need it
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

// go run main.go -> program run

func main() {

	//var conferenceName string = "Go Conference"

	// var bookings  = []string{}  alternative slice
	// bookings := []string{}  alternative slice

	greetUsers()

	// Loop
	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidEmail && isValidTicketNumber && isValidName {

			bookTicket(userTicket, firstName, lastName, email)

			// fmt.Printf("The whole slice : %v\n", bookings)
			// fmt.Printf("The first value of slice : %v\n", bookings[0])
			// fmt.Printf("Type of slice: %T\n", bookings)
			// fmt.Printf("slice size %v\n", len(bookings))

			// Print first name
			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			// discover mistakes at compile time,  not at runtime
			// %T get data type
			// fmt.Printf("User %T type, %T ticket type\n", firstName, userTicket)

			// uint -> whole positive int numbers
			// var noTicketRemaining bool = remainingTickets == 0
			// noTicketRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				// end
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			// fmt.Printf("Your input data is invalid. Try again\n")
			// continue // next itaration of loop

			if !isValidEmail {
				fmt.Println("Email adress doesnt contains @ sign")
			}

			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets are %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//splits the string with white space as seperator -> return slices
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adress: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
