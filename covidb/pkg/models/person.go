package models

import (
	"fmt"

	"github.com/google/uuid"
)

// Person represents a person
type Person struct {
	ID            uuid.UUID
	Gender        string
	Age           int
	Area          string
	City          string
	Relationships Relationships
}

// People represents a many person
type People map[uuid.UUID]Person

func (person Person) String() string {
	return fmt.Sprintf("ID: %s, Gender: %s, Age: %d, Area: %s, City: %s",
		person.ID.String(), person.Gender, person.Age, person.Area, person.City)
}
