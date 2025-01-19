package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Relationship struct {
	id                   uuid.UUID
	sourcePeopleID       uuid.UUID
	targetPeopleID       uuid.UUID
	relationshipStrength RelationshipStrength // example: PM <-> EM: Strong, Sales <-> EM: Weak
	relationshipType     RelationshipType
	notes                string
	dependencyLevel      DependencyLevel
	lastUpdate           time.Time
}

func NewRelationship(sourcePeople People, targetPeople People, strength RelationshipStrength, relationshipType RelationshipType, dependencyLevel DependencyLevel, notes string) Relationship {
	return Relationship{
		id:                   uuid.New(),
		sourcePeopleID:       sourcePeople.id,
		targetPeopleID:       targetPeople.id,
		relationshipStrength: strength,
		relationshipType:     relationshipType,
		notes:                notes,
		dependencyLevel:      dependencyLevel,
		lastUpdate:           time.Now(),
	}
}

func (r Relationship) String() string {
	return fmt.Sprintf("%s -> %s (%s)", r.sourcePeopleID, r.targetPeopleID, r.relationshipStrength)
}
