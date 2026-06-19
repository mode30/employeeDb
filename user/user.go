package user

import (
	// "fmt"
	"errors"
	"fmt"
	// "example.com/employee_management/employee_shared"
)

// type User struct {
// 	FirstName string
// 	LastName  string
// 	Age       int8
// }
type User struct {
	firstName string
	lastName  string
	age       int8
}

func NewUser(firstname string, lastname string, age int8) (*User, error) {
	if firstname == "" || lastname == "" {
		return nil, errors.New("cannot leave first name and last name empyty")
	}
	return &User{

		firstName: firstname,
		lastName:  lastname,
		age:       age,
		// FirstName: firstname,
		// LastName:  lastname,
		// Age:       age,
	}, nil
}


func (person *User)GetUserFirstName()string{
	// return person.FirstName
	return person.firstName
}


func (person *User)GetUserLastName()string{
	return person.lastName
	// return person.LastName
}

func (person *User)GetUserAge()int8{
	// return person.Age
	return person.age
}

func (person *User)DisplayInformation(){
	personal_information:=fmt.Sprintln(person.firstName,person.lastName,person.age)
	fmt.Println(personal_information)
}
