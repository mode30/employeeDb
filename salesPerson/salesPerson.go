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

func (salesPerson *SalesPerson) GetSalesPersonFirstName() string {
firstName:=salesPerson.GetUserFirstName()
	return firstName
	// result := fmt.Sprintln(salesPerson.ReturnFirstName)
	// return result
}

func (salesPerson *SalesPerson) GetSalesPersonLastName() string {

lastName:=salesPerson.GetUserLastName()
	return lastName
	// result := fmt.Sprintln(salesPerson.ReturnLastName)
	// return result
}

func (salesPerson *SalesPerson) GetSalesPersonAge() int8 {
	return salesPerson.GetUserAge()
}

func (salesPerson *SalesPerson) GetId() uint32 {
	return salesPerson.id
}
func (salesPerson *SalesPerson) CalculatePay() float64 {
	return salesPerson.hourlyRate * 24
}
func (salesPerson *SalesPerson) GetRoles() string {
	return salesPerson.roles
}

func (salesPerson *SalesPerson) FullName() {
	firstname := salesPerson.GetUserFirstName()
	lastname := salesPerson.GetUserLastName()
	fullName := fmt.Sprintln(firstname + "" + lastname)
	fmt.Println("full name:", fullName)
}

func (salesPerson *SalesPerson) DisplayInformation() {
	fmt.Println(salesPerson.GetUserFirstName(), salesPerson.GetUserLastName(), salesPerson.GetUserAge(), salesPerson.id, salesPerson.roles, salesPerson.hourlyRate, salesPerson.department)
}
