package main

import "fmt"

type Employee struct {
	Name string
	EmpID int
}

func main() {
	new_employee := Employee{"madhuri", 1234}
	fmt.Println(new_employee)
	fmt.Println(new_employee.Name)
	fmt.Println(new_employee.EmpID)
}
