package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// // TwitterHandle is a type alias
// type TwitterHandle = string

// TwitterHandle is a type definition
type TwitterHandle string

// RedirectURL returns a URL to the twitter handle's page
func (th TwitterHandle) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

// Identifiable is an interface for types that implement ID
type Identifiable interface {
	ID() string
}

// Citizen implements Identifiable, for the ID and includes the related
// country
type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

// NewSocialSecurityNumber contstructs a USA Citizen
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

// NewEuropeanUnionIdentifier constructs a Citizen with given country
func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	case europeanUnionIdentifier:
		return v
	case Person:
		euID, _ := v.Citizen.(europeanUnionIdentifier)
		return euID
	default:
		panic("using an invalied type to initalize EU ID")
	}

}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

// Name has a first and last
type Name struct {
	first string
	last  string
}

// FullName returns the Person's First and Last name
func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

// // Employee has a Name
// type Employee struct {
// 	name Name
// }

// Person has a Name and TwitterHandle
type Person struct {
	Name // Embed struct
	// first         string // Duplicated by embedded name
	twitterHandle TwitterHandle
	Identifiable  // Embed interface
	Citizen
}

// NewPerson creates a new Person
func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

// ID returns a Person's ID
func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

// SetTwitterHandle updates the Person's twitter handle
func (p *Person) SetTwitterHandle(handler TwitterHandle) error {
	if len(handler) == 0 {
		p.twitterHandle = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handle must start with an @ symbol")
	}
	p.twitterHandle = handler
	return nil
}

// TwitterHandle returns the Person's twitter handle
func (p *Person) TwitterHandle() TwitterHandle {
	return p.twitterHandle
}
