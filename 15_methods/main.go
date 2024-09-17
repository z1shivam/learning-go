package main

import "fmt"

func main() {
	fmt.Println("methods in go")
	shivam := User{"shivam kumar", "shivam@mail.com", 18, true}

	shivam.GetStatus()
	shivam.SetNewEmail()
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("is user active", u.Status)
}

func (u User) SetNewEmail() {
	u.Email = "anewemail@gmial.com"
	fmt.Println(u.Name, "'s new email is:", u.Email)
}
