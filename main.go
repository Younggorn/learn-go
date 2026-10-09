package main

import "fmt"

// func increaseSalary(salary *int) {
// 	*salary += 5000
// }

func (e *Employee) increaseSalary(amount int) {
	e.Salary += amount
}

// func calculateOT(otHours, otRate, salary int) (int, int) {
// 	ot := otHours * otRate
// 	total := salary + ot
// 	return ot, total
// }

func (e Employee) calculateSalary() (int, int) {
	ot := e.OtHours * e.OtRate
	total := e.Salary + ot
	return ot, total
}

type Employee struct {
	Name    string
	Salary  int
	OtHours int
	OtRate  int
}

func main() {
	employee := Employee{
		Name:    "Phankorn",
		Salary:  30000,
		OtHours: 10,
		OtRate:  150,
	}
	fmt.Println("Old Salary:", employee.Salary)
	employee.increaseSalary(5000)
	fmt.Println("New Salary:", employee.Salary)
	ot, total := employee.calculateSalary()
	fmt.Println("OT:", ot)
	fmt.Println("Total:", total)

}
