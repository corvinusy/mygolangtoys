package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

	fmt.Println(string(hash))
	return

	hash = []byte("$2y$10$cPvQsLBZfDpEjVgTC2du0OInykzk8GQSlB5YbFcfDnaPdFERdK4y2")

	if bcrypt.CompareHashAndPassword(hash, []byte("albert@gmail.com")) == nil {
		fmt.Println("Password matches")
	} else {
		fmt.Println("Password incorrect!")
	}

	return

}
