package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// can only contain functions
type Identifiable interface {
	ID() string
}

type Citizen interface {
	Country() string
	Identifiable
}

/* type Conflict interface {
	ID() string
} */

type socialSecurityNumber string

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
	default:
		panic("using invalid type to initialize EU Identifier")
	}

}

func (eui europeanUnionIdentifier) ID() string {
	return string(eui.id)
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", string(eui.country))
}

// Test
type Handle struct {
	handle string
	name   string
}

type TwitterHandler Handle

// End Test

// type alias (exact type)
// type TwitterHandle = string

// type declaration (only copies fields)
type TwitterHandle string

func (th TwitterHandle) RedirectUrl() string {
	cleanHandle := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandle)
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	name Name
}

type Person struct {
	Name
	first         string
	last          string
	twitterHandle TwitterHandle
	// Identifiable  // interface can be embeded into a struct
	// identifiable Identifiable
	Citizen
}

func (p Person) ID() string {
	// return "12345"
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

// Embedded types.
func (p *Person) string() string {
	return "I am John Doe"
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) SetTwitterHandle(handle TwitterHandle) error {
	if len(handle) == 0 {
		p.twitterHandle = handle
	} else if !strings.HasPrefix(string(handle), "@") {
		return errors.New("twitter hanle must start with an @ symbol")
	}

	p.twitterHandle = handle
	return nil
}

func (p *Person) TwitterHandle() TwitterHandle {
	return p.twitterHandle
}
