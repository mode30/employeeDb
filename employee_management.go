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
	person_2,err:=user.NewUser("max","tennyson" , 24)
	if err !=nil{
		fmt.Println(err)
		return
	}

	person_3,err:=user.NewUser("kevin","levin" , 14)
	if err !=nil{
		fmt.Println(err)
		return
	}
	person_4,err:=user.NewUser("max","uzumaki" , 16)
	if err !=nil{
		fmt.Println(err)
		return
	}
	person_5,err:=user.NewUser("loligag","pension" , 44)
	if err !=nil{
		fmt.Println(err)
		return
	}

	engineer_1,err:=engineer.NewEngineer(1, 23,  "junior engineer" ,department.Electronics, person_1)
	if err !=nil{
		fmt.Println(err)
		return
	}

	engineer_2,err:=engineer.NewEngineer(1, 23,  "intermediate engineer" ,department.Electronics, person_2)
	if err != nil{
		fmt.Println(err)
		return
	}

	engineer_3,err:=engineer.NewEngineer(1, 23,  "intermediate engineer" ,department.Electronics, person_2)
	if err != nil{
		fmt.Println(err)
		return
	}
	engineer_4,err:=engineer.NewEngineer(1, 23,  "intermediate engineer" ,department.Electronics, person_2)
	if err != nil{
		fmt.Println(err)
		return
	}
	engineer_5,err:=engineer.NewEngineer(1, 23,  "intermediate engineer" ,department.Electronics, person_2)
	if err != nil{
		fmt.Println(err)
		return
	}
	engineer_1.DisplayEngineerInformation()
	engineer_2.DisplayEngineerInformation()
	engineer_3.DisplayEngineerInformation()
	engineer_4.DisplayEngineerInformation()
	engineer_5.DisplayEngineerInformation()



	person_1.DisplayInformation()
	person_2.DisplayInformation()
	person_3.DisplayInformation()
	person_4.DisplayInformation()
	person_5.DisplayInformation()



	}
