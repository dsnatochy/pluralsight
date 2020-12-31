package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	// p := organization.Person{}
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))

	err := p.SetTwitterHandle("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandle("test")) // %T - type, prints "string"
	fmt.Println(p)
	if err != nil {
		fmt.Printf("An error occurred setting twitter handle: %s\n", err.Error())
	}
	/*
		println(p.TwitterHandle().RedirectUrl())
		fmt.Println(p.ID())
		fmt.Println(p.Country())
		// fmt.Println(p.Name.FullName()) - equivalent as below
		fmt.Println(p.FullName())
	*/

	// as long as the Go-type has a predictable memory layout
	// you get equality and hashibility for free

	// name1 := Name{First: "James", Last: "Wilson"}
	// name2 := Name{First: "James", Last: "Wilson"}

	// // this does not work if the struct has
	// // either a slice, map or func member
	// if name1 == name2 {
	// 	println("We match")
	// }

	// name3 := Name{First: "James", Last: "Wilson"}
	// name4 := OtherName{First: "James", Last: "Wilson"}

	// this comparison is not allowed
	// if name3 == name4 {

	// comparing of interfaces is allowed
	ssn := organization.NewSocialSecurityNumber("123-45-5567")
	eu := organization.NewEuropeanUnionIdentifier(12345, "France")
	eu2 := organization.NewEuropeanUnionIdentifier("12345", "France")

	// eu == eu2 - true

	if ssn == eu2 {
		println("we match")
	}

	fmt.Printf("%T\n", ssn)
	fmt.Printf("%T\n", eu)

	// Zero Value Comparison.
	name1 := Name{"", ""}
	if name1 == (Name{}) {
		println("we match") // true
	}

	name2 := &Name{First: "", Last: ""}
	name2 = nil
	if name2 == nil {
		println("name2 is nil")
	}

	// Using struct as a key in a map
	portfolio := map[Name][]organization.Person{}
	portfolio[name1] = []organization.Person{p}

	// custom compare using Equals method
	name3 := Name2{First: "John", Last: "Smith", Middle: []string{"Kelly"}}
	name4 := Name2{First: "John", Last: "Smith", Middle: []string{"Kelly"}}
	if name3.Equals(name4) {
		println("name3 and name4 are equal")
	}

	// Switch
	println(eu.ID())
}

type Name struct {
	First string
	Last  string
}

type OtherName struct {
	First string
	Last  string
}

type Name2 struct {
	First  string
	Last   string
	Middle []string
}

func (n1 Name2) Equals(n2 Name2) bool {
	return n1.First == n2.First &&
		n1.Last == n2.Last &&
		n1.Middle[0] == n2.Middle[0] // unsafe, using for illustration purposes only
}
