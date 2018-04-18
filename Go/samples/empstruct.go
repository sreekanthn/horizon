package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var emp Employee
	emp.Name = "Sreekanth"
	emp.Position = "Manager"
	fmt.Println(emp.Position)

	promoteemp(emp)
	fmt.Println(emp.Position) // Check if the position chnged from manager - passed by value - no change in calling function

	promoteempptr(&emp)
	fmt.Println(emp.Position) // Observe change in position !

	// literal instantiation
	emp1 := Employee{Name: "Randy", Position: "Director"}
	fmt.Println(emp1.Position)
	emp2 := Employee{Name: "Sreekanth", Position: "Senior Manager"}
	// compare the struct instances
	fmt.Println(emp == emp1)
	fmt.Println(emp == emp2)
}

func promoteemp(emp Employee) {
	emp.Position = "Senior " + emp.Position
}

func promoteempptr(emp *Employee) {
	emp.Position = "Senior " + emp.Position
}
