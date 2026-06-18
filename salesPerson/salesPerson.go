package main

import (
	"fmt"
)


type SalesPerson struct {
	id         uint32
	hourlyRate float64
	roles      string
	department DEPARTMENT
	User
}


// -----------------------------------------------------
// salesPerson
func newSalesPerson(id uint32,hourlyRate float64,roles string,user_department DEPARTMENT,user *User)(*SalesPerson,error){
	return &SalesPerson{
		id:id,
		hourlyRate: hourlyRate,
		roles:roles,
		department:user_department,
		user:*user,
	},nil

}
func (salesPerson *SalesPerson) SetfirstName(firstname string) string {

	salesPerson.user.firstName=firstname
	return salesPerson.user.firstName
}

func (salesPerson *SalesPerson) SetLastName(lastname string) string {

	salesPerson.user.lastName=lastname
	return salesPerson.user.lastName

}


func (salesPerson *SalesPerson)displayInformationEngineer(){
	fmt.Println(salesPerson.user.firstName,engineer.user.lastName,engineer.id,engineer.roles,engineer.department,engineer.hourlyRate)
}
