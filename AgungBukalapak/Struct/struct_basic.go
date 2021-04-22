package main

import "fmt"

type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
	Email 		string
	IsActive 	bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "sandi"
	user.LastName = "permana"
	user.Email ="sandipermana88@gmail.com"
	user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "asti"
	user2.LastName = "putri"
	user2.Email ="putri90@gmail.com"
	user2.IsActive = false

	user3 := User{
		ID : 3,
		FirstName : "nana",
		LastName : "aisyah",
		Email :"ginaaisyah@gmail.com",
		IsActive : true,

	}

	user4 := User{4, "mehmed", "abdurahman", "mehmed@gmail.com", false }

	fmt.Println(user)
	fmt.Println(user2.Email)
	fmt.Println(user3.FirstName)
	fmt.Println(user4.LastName)
}