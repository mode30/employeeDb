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
	db.employeedb[id]=employeeInfo
	return db
	// fullName:=employeeInfo.FullName()
}


func(db *EmployeeDB)SearchById(id int16)(value employee_shared.Employee,bool){

	value,exists:=db.employeedb[id]
	return value,exists
	// if exists{
	// 	return EmployeeDB, false
	// }
}
func (db *EmployeeDB)DeleteById(id int16)(value employee_shared.Employee,string){
	value,exists:=db.employeedb[id]
	if exists{
		delete(employee_shared.Employee,id)
	}
	not_found:=fmt.Sprintln("id not found, entry not deleted")
	return not_found

}
