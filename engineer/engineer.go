package engineer

import (
	"example.com/employee_management/department"
	"example.com/employee_management/user"
	"fmt"
)

type Engineer struct {
	id         uint32
	hourlyRate float64
	roles      string
	department department.DEPARTMENT
	user       user.User
}

// -----------------------------------------------------
// Engineer
func newEngineer(id uint32, hourlyRate float64, roles string, user_department department.DEPARTMENT, user *user.User) (*Engineer, error) {
	return &Engineer{
		id:         id,
		hourlyRate: hourlyRate,
		roles:      roles,
		department: user_department,
		user:       *user,
	}, nil

}
func (engineer *Engineer) GetFirstName() string {
	firstName := engineer.user.GetUserFirstName()
	return firstName
	// lastName := engineer.ReturnFirstName()
	// return lastName
}

func (engineer *Engineer) GetLastName() string {
	lastName := engineer.user.GetUserLastName()
	return lastName
}

func (engineer *Engineer) GetAge() int8 {
	return engineer.user.GetUserAge()
}
func (engineer *Engineer) GetId() uint32 {
	return engineer.id

}

func (engineer Engineer) CalculatePay() float64 {
	return engineer.hourlyRate * 24
}
func (engineer Engineer) GetRoles() string {
	return engineer.roles
}

func (engineer *Engineer) FullName() string {
	firstName := engineer.user.GetUserFirstName()
	lastName := engineer.user.GetUserLastName()
	fullName := firstName + "" + lastName
	return fullName
}
func (engineer *Engineer) displayInformationEngineer() {
	fmt.Println(engineer.user.GetUserFirstName(), engineer.user.GetUserLastName(), engineer.user.GetUserAge(), engineer.id, engineer.hourlyRate, engineer.roles, engineer.department)
	// fmt.Println(engineer.user.firstName, engineer.user.lastName, engineer.id, engineer.roles, engineer.department, engineer.hourlyRate)
}
