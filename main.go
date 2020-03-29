package main

import (
	"fmt"

	"github.com/joshuaswickirl/person-object/organization"
)

func main() {
	p := organization.NewPerson("Joshua", "Swick",
		organization.NewEuropeanUnionIdentifier(123456789, "Germany"))
	err := p.SetTwitterHandle("@joshuaswickirl")
	// fmt.Printf("%T\n", organization.TwitterHandle("test"))
	if err != nil {
		fmt.Printf("an error occurred setting twitter handler: %s\n",
			err.Error())
	}

	println(p.ID())
	// println(p.FullName())
	// println(p.TwitterHandle())
	// println(p.TwitterHandle().RedirectURL())
	// println(p.Country())

	name1 := Name{first: "", last: ""}
	name2 := Name{first: "James", last: "Wilson"}
	name3 := Name{first: "James", last: "Wilson"}

	if name2 == name3 && &name2 != &name3 {
		println("name 2 and 3 match values, but are different" +
			"instances of the Name struct")
	}

	// ssn := organization.NewSocialSecurityNumber("123-45-6789")
	// eu := organization.NewEuropeanUnionIdentifier("12345", "France")
	// eu2 := organization.NewEuropeanUnionIdentifier("12345", "France")
	portfolio := map[Name][]organization.Person{}
	portfolio[name1] = []organization.Person{p}

	person := portfolio[name1][0]
	println(person.Name.FullName())

	if name1 == (Name{}) {
		println("We match")
	}

	// fmt.Printf("%T\n", ssn)
	// fmt.Printf("%T\n", eu)

}

// Name has a first and last
type Name struct {
	first string
	last  string
}
