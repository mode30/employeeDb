package department

import "fmt"

// "fmt"

type DEPARTMENT string

const (
	Electronics DEPARTMENT = "Electronics"
	Engineering DEPARTMENT = "Engineering"
	Sales       DEPARTMENT = "Sales"
	Marketing   DEPARTMENT = "Marketing"
)

type DepartmentDisplayer interface {
	DisplayDepartments()
}

var AllDepartment = []DEPARTMENT{"Electronics", "Engineering", "Sales", "Marketing"}

func ListAllDepartment(department *DEPARTMENT) {
	for _, d := range *department {
		fmt.Println(d)
	}

}

// func (department *DEPARTMENT)DisplayDepartment(){
// for
//
// 	for departments:=range DEPARTMENT{
// 		fmt.Println(a ...any)
// 	}
