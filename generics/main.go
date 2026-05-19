package main

import "fmt"

type User struct {
	Email  string `json:"email"`
	Number int64  `json:"number"`
}

var db = make(map[User]bool)

func signInValid[T string | int64](inp T) string {

	switch v := any(inp).(type) {
	case string:
		for user := range db {
			if user.Email == v {
				return "User Already exists"
			}
		}
	case int64:
		for user := range db {
			if user.Number == v {
				return "User Already exists"
			}
		}
	}
	return "Please Sign In"
}

func main() {
	andy := User{Email: "andy.05@gmail.com", Number: 2356705}
	res := signInValid(andy.Email)
	fmt.Printf("%s", res)
}
