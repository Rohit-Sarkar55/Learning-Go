package main

import "fmt"

type Payable interface {
	CalculatePay() float64
	fmt.Stringer
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried Employee [Name : %s , Salary: $%g]", se.Name, se.CalculatePay())
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly Employee[Name: %s , Salary: $%g]", he.Name, he.CalculatePay())
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	SalesAmount    float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + ce.CommissionRate*ce.SalesAmount
}
func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission Employee [Name %s , Salary $%g", ce.Name, ce.CalculatePay())
}

func printEmpSummary[T fmt.Stringer](emp T) {
	fmt.Printf(" - Processing %s\n", emp)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("Processing payroll...")

	totalMonthlySalary := 0.0
	for _, employee := range employees {
		printEmpSummary(employee)

		pay := employee.CalculatePay()
		totalMonthlySalary += pay
		fmt.Printf("		Monthly Pay: %.2f\n", pay)
	}

	fmt.Printf("Total Monthly Payroll: %.2f\n_____________________________\n", totalMonthlySalary)

}

func main() {

	se := SalariedEmployee{Name: "Rohit", AnnualSalary: 124526.0}

	he := HourlyEmployee{Name: "Rahul", HourlyRate: 52.56, HoursWorked: 165}

	ce := CommissionEmployee{Name: "Shreya", BaseSalary: 6700, CommissionRate: 2.67, SalesAmount: 567}

	paylist := []Payable{se, he, ce, HourlyEmployee{"Diya", 36.34, 170}}

	ProcessPayroll(paylist)

}
