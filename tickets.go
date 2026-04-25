package main

import (
	"fmt"
	"time"
)

var hash = make(map[int]bool)

type Ticket struct {
	SeatNo    int
	movieName string
	NoTickets int32
	isBooked  bool
	BookTime  time.Time
}

// Map to store all tickets
var tickets = make(map[int]*Ticket)

func (t *Ticket) bookTicket() bool {
	_, ok := hash[t.SeatNo]
	if ok {
		fmt.Println("Seat Already Taken!")
		return false
	}
	hash[t.SeatNo] = true
	t.isBooked = true
	t.BookTime = time.Now()
	tickets[t.SeatNo] = t
	fmt.Printf("Ticket booked successfully for Seat %d - Movie: %s\n", t.SeatNo, t.movieName)
	return true
}

func (t *Ticket) cancelTicket() bool {
	_, ok := hash[t.SeatNo]
	if !ok {
		fmt.Println("Seat not booked!")
		return false
	}
	delete(hash, t.SeatNo)
	delete(tickets, t.SeatNo)
	t.isBooked = false
	fmt.Printf("Ticket cancelled for Seat %d\n", t.SeatNo)
	return true
}

func isAvailable(seatNo int) bool {
	_, ok := hash[seatNo]
	return !ok
}

func displayTickets() {
	if len(tickets) == 0 {
		fmt.Println("No tickets booked!")
		return
	}
	fmt.Println("\n========== Booked Tickets ==========")
	for seatNo, ticket := range tickets {
		fmt.Printf("Seat %d | Movie: %s | Tickets: %d | Booked At: %s\n",
			seatNo, ticket.movieName, ticket.NoTickets, ticket.BookTime.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("====================================\n")
}

func TicketManager() {
	fmt.Println("===== Movie Ticket Management =====")

	// Create and book tickets
	ticket1 := &Ticket{
		SeatNo:    1,
		movieName: "Inception",
		NoTickets: 2,
	}
	ticket1.bookTicket()

	ticket2 := &Ticket{
		SeatNo:    2,
		movieName: "The Matrix",
		NoTickets: 1,
	}
	ticket2.bookTicket()

	// Try booking same seat
	ticket3 := &Ticket{
		SeatNo:    1,
		movieName: "Inception",
		NoTickets: 1,
	}
	ticket3.bookTicket()

	displayTickets()
	fmt.Printf("Seat 3 available: %v\n", isAvailable(3))
	fmt.Printf("Seat 1 available: %v\n", isAvailable(1))

	ticket1.cancelTicket()
	displayTickets()
}
