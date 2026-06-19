package salesPerson

import (
	"fmt"

	"example.com/employee_management/department"
	// "example.com/employee_management/engineer"
	// "example.com/employee_management/salesPerson"

	// "example.com/employee_management/salesPerson"
	"example.com/employee_management/user"
)


type SalesPerson struct {
	id         uint32
	hourlyRate float64
	roles      string
	department department.DEPARTMENT
	user.User
}

// -----------------------------------------------------
// salesPerson
//returns a new cosntructor for newsales perseron
func newSalesPerson(id uint32, hourlyRate float64, roles string, user_department department.DEPARTMENT, user_entry *user.User) (*SalesPerson, error) {
	return &SalesPerson{
		id:         id,
		hourlyRate: hourlyRate,
		roles:      roles,
		department: user_department,
		User:       *user_entry,
	}, nil

}

//returns first name of sales person
func (salesPerson *SalesPerson) GetSalesPersonFirstName() string {
firstName:=salesPerson.GetUserFirstName()
	return firstName
	// result := fmt.Sprintln(salesPerson.ReturnFirstName)
	// return result
}

//retruns last name of sales person
func (salesPerson *SalesPerson) GetSalesPersonLastName() string {

lastName:=salesPerson.GetUserLastName()
	return lastName
	// result := fmt.Sprintln(salesPerson.ReturnLastName)
	// return result
}

//returns age of sales eprson
func (salesPerson *SalesPerson) GetSalesPersonAge() int8 {
	return salesPerson.GetUserAge()
}

//returns id of sales eprson
func (salesPerson *SalesPerson) GetId() uint32 {
	return salesPerson.id
}
//returns salary of sales peerson
func (salesPerson *SalesPerson) CalculatePay() float64 {
	return salesPerson.hourlyRate * 24
}
//returns roles of sales person(postion)
func (salesPerson *SalesPerson) GetRoles() string {
	return salesPerson.roles
}

//returns full name of sales person
func (salesPerson *SalesPerson) FullName() {
	firstname := salesPerson.GetUserFirstName()
	lastname := salesPerson.GetUserLastName()
	fullName := fmt.Sprintln(firstname + "" + lastname)
	fmt.Println("full name:", fullName)
}

//displays fullinformation of sales person
func (salesPerson *SalesPerson) DisplayInformation() {
	fmt.Println(salesPerson.GetUserFirstName(), salesPerson.GetUserLastName(), salesPerson.GetUserAge(), salesPerson.id, salesPerson.roles, salesPerson.hourlyRate, salesPerson.department)
}
