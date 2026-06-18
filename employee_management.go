package main

import (
	"fmt"
	"errors"
)



type DEPARTMENT string

const (
	Electronics DEPARTMENT = "Electronics"
	Engineering   DEPARTMENT = "Engineering"
	Sales       DEPARTMENT = "Sales"
	Marketing  DEPARTMENT="Marketing"
)
type User struct {
	firstName string
	lastName  string
	age       int8
}
type Employee interface {
	ReturnFirstName() string
	ReturnLastName() string
	ReturnAgeName() int8
	GetId() uint32
	CalculatePay() float64
	GetRoles() string
	SetfirstName(firstname string) string
	SetLastName(lastname string) string
	FullName() string
	DisplayInformation()
}
type Displayer interface{
	Display()
}

func main() {
	new_engineer,err:=newUser("benjamin",  "carson", 33)
	if err !=nil{
		return
	}

	// new_engineer.displ


}


func newUser(firstname string,lastname string, age int8)(*User,error){
	if firstname=="" || lastname ==""{
		return nil,errors.New("cannot leave first name and last name empyty")
	}
	return &User{
		firstName:firstname,
		lastName:lastname,
		age:age,
	},nil
}
