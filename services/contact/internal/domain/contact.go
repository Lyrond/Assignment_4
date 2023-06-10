package domain

type Contact struct {
	Id       int
	fullName fullName
	Phone    int
}

type fullName struct {
	firstName  string
	secondName string
	middleName string
}

var contactCounter int

func (c Contact) GetFullName() fullName {
	return c.fullName
}

func NewContact(firstName, secondName, middleName string, phone int) *Contact {
	contactCounter++
	name := &fullName{
		firstName:  firstName,
		secondName: secondName,
		middleName: middleName,
	}

	return &Contact{
		Id:       contactCounter,
		fullName: *name,
		Phone:    phone,
	}
}
