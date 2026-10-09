package main

import "fmt"

type Employee struct {
	Name    string
	Salary  int
	OtHours int
	OtRate  int
}

func (e *Employee) increaseSalary(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("salary increase must be greater than zero")
	}

	e.Salary += amount
	return nil
}

func (e Employee) calculateSalary() (int, int) {
	ot := e.OtHours * e.OtRate
	total := e.Salary + ot

	return ot, total
}

f

func main() {
	users := []Employee{
		{Name: "John Doe", Salary: 50000, OtHours: 10, OtRate: 20},
		{Name: "Jane Smith", Salary: 60000, OtHours: 5, OtRate: 25},
		{Name: "Bob Johnson", Salary: 55000, OtHours: 8, OtRate: 22},
	}

	for i := range users {
		fmt.Println("Name:", users[i].Name)
		fmt.Println("Old Salary:", users[i].Salary)

		err := users[i].increaseSalary(-100)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("New Salary:", users[i].Salary)

		ot, total := users[i].calculateSalary()

		fmt.Println("OT:", ot)
		fmt.Println("Total:", total)
		fmt.Println()
	}

	
}