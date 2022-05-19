package database

import (
	"fmt"
	"html/template"

	"golang.org/x/crypto/bcrypt"
)

//This File Provide the necessary functions to carry out
//Initialization,adding, deleting of customer data needed from the handler functions
type User struct {
	CustomerId int
	Username   string
	Password   []byte
	Firstname  string
	Lastname   string
	IsAdmin    bool
	BookingId  []int
}

var tpl *template.Template
var Users = map[string]*User{}
var Sessions = map[string]string{}

func InitCustomers() {
	list := []*User{
		{
			Username: "admin",
			Password: []byte("1234"),
			IsAdmin:  true,
		}, {
			Username:  "user",
			Firstname: "John",
			Lastname:  "Deo",
			Password:  []byte("1234"),
			IsAdmin:   false,
			BookingId: []int{1, 2, 3, 4},
		},
	}

	for _, value := range list {
		CreateNewUser(value)
	}
}

func CustomerId() int {
	max := 0
	for _, value := range Users {
		if value.CustomerId > max {
			max = value.CustomerId
		}
	}
	return max + 1
}
func CreateNewUser(u *User) error {
	u.CustomerId = CustomerId()

	bpassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("CreateNewUser: %w", err)
	}
	u.Password = bpassword
	Users[u.Username] = u
	return nil

}
func InitilizeUsers() {
	list := []*User{
		{
			Username: "admin",
			Password: []byte("1234"),
			IsAdmin:  true,
		}, {
			Username:  "user",
			Firstname: "John",
			Lastname:  "Doe",
			Password:  []byte("1234"),
			IsAdmin:   false,
			BookingId: []int{1, 2, 3},
		},
	}

	for _, u := range list {
		CreateNewUser(u)
	}
}

func (u *User) CancelBookings(id int) error {
	for result, value := range u.BookingId {
		if value == id {
			u.BookingId = append(u.BookingId[:result], u.BookingId[result+1:]...)
			return nil
		}
	}
	return fmt.Errorf("booking ID was not found")
}
