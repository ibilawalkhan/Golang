package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // Waits for the launched goroutine to finish

func main() {
	conferenceName := "Go Conference" // This syntax apply to only variables not constants, this declare and assign type at the same time
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = make([]UserData, 0)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}

			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)

			wg.Add(1) 
			go sendTicket(userTickets, firstName, lastName, email) // Goroutine

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("Remaining tickets %v\n", remainingTickets)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name you entered is too short")
			} else if !isValidEmail {
				fmt.Println("Your email address does not contain @")
			} else if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers(conferenceName string, confTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to our booking application %v\n", conferenceName)
	fmt.Println("We have total of", confTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket :%v\n to email address %v\n", ticket, email)
	fmt.Println("\n###################")
	wg.Done() // Tells the WaitGroup that the goroutine is done
}
