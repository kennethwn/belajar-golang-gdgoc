package main

import "fmt"

type UserCity struct {
	name, city string
}

type User struct {
	id       int
	name     string
	username string
	password string
}

func (u User) getName() string {
	return u.name
}

func main() {
	/* Struct without pointer */
	user := UserCity{"Kenneth", "Jakarta"}
	newUser := user

	user.city = "Bandung"

	fmt.Println(user)
	fmt.Println(newUser)

	/* Declare, assign, accessing struct */
	var user1 User
	user1.id = 1
	user1.name = "Kenneth William"
	user1.username = "kennethwilliam"
	user1.password = "backend2025"

	user2 := User{
		id:       2,
		name:     "Joshua",
		username: "joshuawenata",
		password: "joshua123",
	}

	fmt.Println(user1)
	fmt.Println(user2)

	fmt.Println(user1.getName())
	fmt.Println(user2.getName())
}
