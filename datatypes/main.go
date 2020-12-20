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

	println(p.TwitterHandle().RedirectUrl())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.FullName())
}
