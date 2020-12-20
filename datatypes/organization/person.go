package organization

import (
	"errors"
	"fmt"
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

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return string(eui.id)
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", string(eui.country))
}

type Name struct {
	first string
	last  string
}

type Employee struct {
	Name // embedding
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

type Person struct {
	Name          // embeding
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

func (n Name) FullName() string {

	return fmt.Sprintf("%s %s", n.first, n.last)
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
