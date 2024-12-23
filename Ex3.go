package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d Tenge", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d Tenge, Hours Worked: %.2f", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{Employees: make(map[string]Employee)}
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.Employees[fmt.Sprint(e.ID)] = emp
	case PartTimeEmployee:
		c.Employees[fmt.Sprint(e.ID)] = emp
	}
	fmt.Println("Employee added successfully.")
}

func (c *Company) ListEmployees() {
	if len(c.Employees) == 0 {
		fmt.Println("No employees to list.")
		return
	}
	for id, emp := range c.Employees {
		fmt.Printf("Employee ID: %s, Details: %s\n", id, emp.GetDetails())
	}
}

func main() {
	c := NewCompany()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Choose an option: 1. Add Employee 2. List Employees 3. Exit")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid option. Please enter a number.")
			continue
		}

		switch option {
		case 1:
			fmt.Println("Enter employee type (full-time/part-time):")
			if !scanner.Scan() {
				break
			}
			typeInput := strings.ToLower(scanner.Text())
			if typeInput == "full-time" {
				fmt.Println("Enter ID, Name, and Salary (space-separated):")
				if !scanner.Scan() {
					break
				}
				fields := strings.Fields(scanner.Text())
				if len(fields) != 3 {
					fmt.Println("Invalid input. Please enter ID, Name, and Salary.")
					continue
				}
				id, err1 := strconv.ParseUint(fields[0], 10, 64)
				salary, err2 := strconv.ParseUint(fields[2], 10, 32)
				if err1 != nil || err2 != nil {
					fmt.Println("Invalid input. Please ensure ID and Salary are numbers.")
					continue
				}
				emp := FullTimeEmployee{
					ID:     id,
					Name:   fields[1],
					Salary: uint32(salary),
				}
				c.AddEmployee(emp)
			} else if typeInput == "part-time" {
				fmt.Println("Enter ID, Name, Hourly Rate, and Hours Worked (space-separated):")
				if !scanner.Scan() {
					break
				}
				fields := strings.Fields(scanner.Text())
				if len(fields) != 4 {
					fmt.Println("Invalid input. Please enter ID, Name, Hourly Rate, and Hours Worked.")
					continue
				}
				id, err1 := strconv.ParseUint(fields[0], 10, 64)
				rate, err2 := strconv.ParseUint(fields[2], 10, 64)
				hours, err3 := strconv.ParseFloat(fields[3], 32)
				if err1 != nil || err2 != nil || err3 != nil {
					fmt.Println("Invalid input. Please ensure ID, Hourly Rate, and Hours Worked are numbers.")
					continue
				}
				emp := PartTimeEmployee{
					ID:          id,
					Name:        fields[1],
					HourlyRate:  rate,
					HoursWorked: float32(hours),
				}
				c.AddEmployee(emp)
			} else {
				fmt.Println("Invalid employee type. Please enter 'full-time' or 'part-time'.")
			}
		case 2:
			c.ListEmployees()
		case 3:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}
	}
}
