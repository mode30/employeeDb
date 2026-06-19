package employee_shared

import (
	"fmt"

	"example.com/employee_management/engineer"
	"example.com/employee_management/salesPerson"




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
	type Saver interface{
		Saver()
	}
	type ottputable interface{
		Saver()
		Displayer()
	}

	//displays information about user
	func DisplayInformation[T engineer.Engineer|SalesPerson](data T){

		fmt.Print("%v\n",data)
	}

	//returns fullname of user
	func FullName()string{
		firstName:=ReturnFirstName()
		LastName:=ReturnLastName()
		fullName:=fmt.Sprintln(firstName +""+LastName)
		return fullName()
	}


	func ReturnFirstName()string{

	}
