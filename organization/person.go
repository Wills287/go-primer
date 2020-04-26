package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Citizen interface {
	Identifiable
	Country() string
}

// Type declaration
type TwitterHandler string

// Type alias
//type TwitterHandler = string

type Identifiable interface {
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Citizen
	Name
	twitterHandler TwitterHandler
}

func NewPerson(citizen Citizen, firstName, lastName string) Person {
	return Person{
		Citizen: citizen,
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

// func (p *Person) ID() string {
// 	return "12345"
// }

func (p *Person) SetTwitterHandler(twitterHandler TwitterHandler) error {
	if len(twitterHandler) == 0 {
		p.twitterHandler = twitterHandler
	} else if !strings.HasPrefix(string(twitterHandler), "@") {
		return errors.New("Twitter Handler must start with an @ symbol")
	}

	p.twitterHandler = twitterHandler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func (twitterHandler TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(twitterHandler), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}
