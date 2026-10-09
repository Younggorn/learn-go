package main

import (
	"employee-app/employee"
	"fmt"
)

func main() {
	users := []employee.Employee{
		{Name: "John Doe", Salary: 50000, OtHours: 10, OtRate: 20},
		{Name: "Jane Smith", Salary: 60000, OtHours: 5, OtRate: 25},
		{Name: "Bob Johnson", Salary: 55000, OtHours: 8, OtRate: 22},
	}

	err := employee.UpdateSalary("Jane Smith", 5000, users)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jane, err := employee.FindEmployee("Jane Smith", users)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", jane.Name)
	fmt.Println("Salary:", jane.Salary)
}