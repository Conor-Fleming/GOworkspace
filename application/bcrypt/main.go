package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//password encryption with bcrypt package
func main() {
	s := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(s)
	//fmt.Println(bs) //hashed version

	login := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println("Wrong password", err)
		return
	}
	fmt.Println("You're logged in")
}
