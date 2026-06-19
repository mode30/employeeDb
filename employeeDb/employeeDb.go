package main

import (
	"fmt"
	"example.com/employee_management/employee_shared"
)
// type employeeDatabase employee_shared.Employee

type EmployeeDB struct{
	employeedb map[int16]employee_shared.Employee
	// employeedb map[int16]employeeDatabase
}

func NewEmployeeDb()*EmployeeDB{
	return &EmployeeDB{
		employeedb: make(map[int16]employee_shared.Employee),
	}
}


func (db *EmployeeDB)AddEmployee(employeeInfo employee_shared.Employee)*EmployeeDB{

	id:=employeeInfo.GetId()
	fullName:=employeeInfo.FullName()
	return &EmployeeDB{
		employeedb:map[id]fullName,
	}
}
