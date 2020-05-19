package models

import "github.com/google/uuid"

// Relationship represents a relationship
type Relationship struct {
	Name     string
	PersonID uuid.UUID
}

// Relationships represents multiple / a list of Relationship
type Relationships []Relationship
