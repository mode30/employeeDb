package main

import (
	// "errors"
	"fmt"

	// "example.com/employee_management/employee_shared"
	"example.com/employee_management/user"
	"example.com/employee_management/department"
	"example.com/employee_management/engineer"
	// "fmt"
)

func main() {
	person_1,err:=user.NewUser("benjamin","carson" , 44)
	if err !=nil{
		fmt.Println(err)
		return
	}
	person_1.DisplayInformation()

	engineer_1,err:=engineer.NewEngineer(1, 23,  "junior engineer" ,department.Electronics, *person_1)
	if err != nil{
		fmt.Println(err)
		return
	}



	}
