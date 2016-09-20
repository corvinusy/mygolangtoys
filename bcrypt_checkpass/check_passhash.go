package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	hash, _ := bcrypt.GenerateFromPassword([]byte("a_test_operator_08"), bcrypt.DefaultCost)

	fmt.Println(string(hash))

	hash = []byte("$2a$10$ububjkFIYwaEdRg6zIohLePQUbP8sfGsXEvO2LX6jKeVRB/weEHTK")

	if bcrypt.CompareHashAndPassword(hash, []byte("operator")) == nil {
		fmt.Println("Password matches")
	} else {
		fmt.Println("Password incorrect!")
	}

	return

}
