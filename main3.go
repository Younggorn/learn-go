package main

import "fmt"

type Employee struct {
	Name    string
	Salary  int
	OtHours int
	OtRate  int
}

// func (e *Employee) increaseSalary(amount int) error {
// 	if amount <= 0 {
// 		return fmt.Errorf("salary increase must be greater than zero")
// 	}

// 	e.Salary += amount
// 	return nil
// }

// func (e Employee) calculateSalary() (int, int) {
// 	ot := e.OtHours * e.OtRate
// 	total := e.Salary + ot

// 	return ot, total
// }



func  findEmployee(name string, users []Employee) (*Employee, error) {
	for i := range users {
		if users[i].Name == name {
			return &users[i], nil
		}
	}

	return nil, fmt.Errorf("employee %q not found" , name)
}

func updateSalary(name string, amount int, users []Employee) error {

	if amount <= 0 {
		return fmt.Errorf("salary increase must be greater than zero")
	}

	employee, err := findEmployee(name, users)
	if err != nil {
		return err
	}

	oldSalary := employee.Salary
	newSalary := oldSalary + amount

	fmt.Println("Name:", employee.Name)
	fmt.Println("Old Salary:", oldSalary)
	fmt.Println("New Salary:", newSalary)
	return nil
}

func main() {
	users := []Employee{
		{Name: "John Doe", Salary: 50000, OtHours: 10, OtRate: 20},
		{Name: "Jane Smith", Salary: 60000, OtHours: 5, OtRate: 25},
		{Name: "Bob Johnson", Salary: 55000, OtHours: 8, OtRate: 22},
	}

	err := updateSalary("Jane Smith", 5000, users)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	

}
