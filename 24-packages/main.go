package main

import (
	"fmt"

	"packaging/auth"
	"packaging/user"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("bob", "secret")

	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "a@b.com",
		Name:  "alice",
	}

	// fmt.Println(user)
	color.Red(user.Email)
	color.Green(user.Name)
}
