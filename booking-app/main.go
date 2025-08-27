package main

// go programs are organized into package
// go standart library, provides different core packages for us to use

// a program can only have 1 main function, because you can only have 1 entypoint
import (
	"booking-app/helper"
	"fmt" // standart Format package
	"sync"

	//"strconv"
	//"strings"
	"time"
)

// packace level variable / we can use normal variable level syntax
// define variable as 'local' as possible
// create the varible where you need it
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = []string{}
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

// empty slice of map , 0 initial size it will increase anyways , initials list of maps
// allow multiple key value pairs per user -> map
// go run main.go -> program run

// type: create a new type
// struct > map ( mixed type )
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//var conferenceName string = "Go Conference"

	// var bookings  = []string{}  alternative slice
	// bookings := []string{}  alternative slice

	greetUsers()

	// Loop

	firstName, lastName, email, userTicket := getUserInput()

	isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

	if isValidEmail && isValidTicketNumber && isValidName {

		bookTicket(userTicket, firstName, lastName, email)

		// Sets the number of goroutines to wait for(increases the counter by the provided number)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email) // go: start a new goroutine
		// goroutine is lightweight thread managed by the Go runtime
		// generating and sending ticket task runs now in the background
		// WE need to tell 'main' that it needs to wait until 'sendTicket' is done

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
			//break
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
	wg.Wait()
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
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
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

	// create a map for user - we can not mix data types
	// var userData = make(map[string]string) //empty map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Handle blocing code -> Goroutunes
// Concurrency to make our program make efficient

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("#############################")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
	fmt.Println("#############################")
	wg.Done()
}
