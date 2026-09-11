package main

import "fmt"

type Contact struct {
	Id    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextId = 1

func init() {

	fmt.Printf("Init method executing.....\n")
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name string, email string, phone string) {
	_, exists := contactIndexByName[name]
	if exists {
		fmt.Printf("contact already exists named %s\n", name)
		return
	}
	newContact := Contact{
		Id:    nextId,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextId++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1

	fmt.Printf("added contact named %v\n", newContact)

}

func ListContacts() {

	if len(contactList) == 0 {
		fmt.Println("No contact found")
		return
	}

	for _, contact := range contactList {
		fmt.Printf("contact %+v\n", contact)
	}
}

func findContactByName(name string) *Contact {
	ind, exits := contactIndexByName[name]
	if !exits {
		fmt.Printf("No contact named %s found \n", name)
		return nil
	}
	return &contactList[ind]
}

func main() {

	ListContacts()
	addContact("Alice", "alice@gmail.com", "12345")
	addContact("Bob", "bob@yahoo.com", "31543")
	addContact("Charlie", "charlie@hotmail.com", "45262")
	addContact("Alice", "alice@gmail.com", "12345")

	ListContacts()

	foundContact := findContactByName("Alice")
	if foundContact != nil {
		fmt.Printf("Found %+v \n", foundContact)
	}
}
