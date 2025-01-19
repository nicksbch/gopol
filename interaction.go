package main

import (
	"time"

	"github.com/google/uuid"
)

type Interaction struct {
	id              uuid.UUID
	date            time.Time
	peopleID        uuid.UUID
	interactionType InteractionType
	sentiment       InteractionSentiment
	impact          InteractionImpact
	notes           string
}

func NewInteraction(people People, interactionType InteractionType, sentiment InteractionSentiment, impact InteractionImpact, notes string) Interaction {
	return Interaction{
		id:              uuid.New(),
		date:            time.Now(),
		peopleID:        people.id,
		interactionType: interactionType,
		sentiment:       sentiment,
		impact:          impact,
		notes:           notes,
	}
}
