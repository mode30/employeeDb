package engineer

import (
	"fmt"
)


type Engineer struct {
	id         uint32
	hourlyRate float64
	roles      string
	department DEPARTMENT
	user User
}


// -----------------------------------------------------
// Engineer
func newEngineer(id uint32,hourlyRate float64,roles string,user_department DEPARTMENT,user *User)(*Engineer,error){
	return &Engineer{
		id:id,
		hourlyRate: hourlyRate,
		roles:roles,
		department:user_department,
		user:*user,
	},nil

}
func (engineer *Engineer) SetfirstName(firstname string) string {

	engineer.user.firstName=firstname
	return engineer.user.firstName
}

func (engineer *Engineer) SetLastName(lastname string) string {

	engineer.user.lastName=lastname
	return engineer.user.lastName

}


func (engineer *Engineer)displayInformationEngineer(){
	fmt.Println(engineer.user.firstName,engineer.user.lastName,engineer.id,engineer.roles,engineer.department,engineer.hourlyRate)
}
