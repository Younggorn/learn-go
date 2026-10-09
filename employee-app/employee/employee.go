package employee

import "fmt"

type Employee struct {
	Name    string
	Salary  int
	OtHours int
	OtRate  int
}

func  FindEmployee(name string, users []Employee) (*Employee, error) {
	for i := range users {
		if users[i].Name == name {
			return &users[i], nil
		}
	}

	return nil, fmt.Errorf("employee %q not found" , name)
}

func UpdateSalary(name string, amount int, users []Employee) error {

	if amount <= 0 {
		return fmt.Errorf("salary increase must be greater than zero")
	}

	employee, err := FindEmployee(name, users)
	if err != nil {
		return err
	}


	employee.Salary += amount
	
	return nil
}
