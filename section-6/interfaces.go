package main

import "fmt"

type Person interface {
	GetName() string
}
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
func main() {

	e1 := Employee{1, "John"}
	b1 := BusinessPerson{1, "Alice"}
	displayPerson(e1)
	displayPerson(b1)

}
