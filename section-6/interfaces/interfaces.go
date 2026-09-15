package main

import "fmt"

type Person interface {
	GetName() string
}

type FL2 float32
type FL3 float32

type Employee struct {
	ID   int
	Name string
}
type BusinessPerson struct {
	ID   int
	Name string
}

func (b BusinessPerson) GetName() string {
	return fmt.Sprintf("%s", b.Name)
}
func (e Employee) GetName() string {
	return fmt.Sprintf("%s", e.Name)
}
func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

// stringer interface

func (b BusinessPerson) String() string {
	return fmt.Sprintf("Business Person [ID: %d , Name: %s ]", b.ID, b.Name)
}
func (b Employee) String() string {
	return fmt.Sprintf("Employee [ID: %d , Name: %s ]", b.ID, b.Name)
}

func (f FL2) String() string {
	return fmt.Sprintf("This is a custome float of 2 decimals %.2f", f)
}

func (f FL3) String() string {
	return fmt.Sprintf("This is a custome float of 3 decimals %.3f", f)
}
func main() {

	e1 := Employee{1, "John"}
	b1 := BusinessPerson{1, "Alice"}
	displayPerson(e1)
	displayPerson(b1)

	fmt.Println(e1)
	fmt.Println(b1)

	//stringer works on primitive data type is well

	f1 := FL2(3.1341531)
	f2 := FL3(3.1341531)

	fmt.Println(f1)
	fmt.Println(f2)

}
