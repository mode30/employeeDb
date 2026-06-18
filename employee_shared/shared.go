package main

import (
	"fmt"

	"example.com/employee_management/engineer"
)


func displyInformation[T engineer.Engineer|SalesPerson](data T){

	// personal_information:=fmt.Sprintf(data)
	// personal_information:=fmt.Sprintf(data)
	fmt.Print("%v\n",data)
	// personal_information:=fmt.Sprintf(engineer.user.firstName,engineer.user.lastName,engineer.user.age,engineer.id,engineer.roles,engineer.hourlyRate)
	// fmt.Println(personal_information)
}
