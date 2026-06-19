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

// returns an instance of engineer
func NewEngineer(id uint32, hourlyRate float64, roles string, user_department department.DEPARTMENT, user *user.User) (*Engineer, error) {
	return &Engineer{
		id:         id,
		hourlyRate: hourlyRate,
		roles:      roles,
		department: user_department,
		user:       *user,
	}, nil

}
//return engineer first name
func (engineer *Engineer) GetFirstName() string {
	firstName := engineer.user.GetUserFirstName()
	return firstName
}

//return engineer lastname
func (engineer *Engineer) GetLastName() string {
	lastName := engineer.user.GetUserLastName()
	return lastName
}

//return engineer age
func (engineer *Engineer) GetAge() int8 {
	return engineer.user.GetUserAge()
}
//return engineer id
func (engineer *Engineer) GetId() uint32 {
	return engineer.id

}

//calcaulate engieer salary
func (engineer Engineer) CalculatePay() float64 {
	return engineer.hourlyRate * 24
}
//get engineer role
func (engineer Engineer) GetRoles() string {
	return engineer.roles
}

//display engineer full name
func (engineer *Engineer) FullName() string {
	firstName := engineer.user.GetUserFirstName()
	lastName := engineer.user.GetUserLastName()
	fullName := firstName + "" + lastName
	return fullName
}
func (engineer *Engineer) displayInformationEngineer() {
	fmt.Println(engineer.user.GetUserFirstName(), engineer.user.GetUserLastName(), engineer.user.GetUserAge(), engineer.id, engineer.hourlyRate, engineer.roles, engineer.department)
}


//display engineers full infromation
func (engineer *Engineer)DisplayEngineerInformation(){
	personal_information:=fmt.Sprintln("information\n:",engineer.user.GetUserFirstName(),engineer.user.GetUserLastName(),engineer.user.GetUserAge(),engineer,engineer.hourlyRate,engineer.id,engineer.roles,engineer.department)
	fmt.Println(personal_information)
}
