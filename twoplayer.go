package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func twoPlayer() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the value for p1: ")
	p1, _ := reader.ReadString('\n')

	fmt.Print("Enter the value for p2: ")
	p2, _ := reader.ReadString('\n')

	p1 = strings.ToLower(strings.TrimSpace(p1))
	p2 = strings.ToLower(strings.TrimSpace(p2))

	switch {
	case p1 == "rocks" && p2 == "scissors",
		p1 == "paper" && p2 == "rocks",
		p1 == "scissors" && p2 == "paper":
		fmt.Println("p1 Wins!")

	case p1 == p2:
		fmt.Println("It's a Draw!")

	case p1 == "rocks" && p2 == "paper",
		p1 == "paper" && p2 == "scissors",
		p1 == "scissors" && p2 == "rocks":
		fmt.Println("p2 Wins!")

	default:
		fmt.Println("Invalid Input!")
	}
}
