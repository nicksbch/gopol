package main

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Organisation struct {
	people        map[uuid.UUID]People
	relationships map[uuid.UUID]Relationship
	interactions  map[uuid.UUID]Interaction
	mutex         sync.RWMutex
}

func NewOrganisation() *Organisation {
	return &Organisation{
		people:        make(map[uuid.UUID]People),
		relationships: make(map[uuid.UUID]Relationship),
		interactions:  make(map[uuid.UUID]Interaction),
	}
}

func (g *Organisation) AddPeople(p People) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.people[p.id] = p
}

func (g *Organisation) AddRelationship(r Relationship) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, exists := g.people[r.sourcePeopleID]; !exists {
		return errors.New("source people not found")
	}
	if _, exists := g.people[r.targetPeopleID]; !exists {
		return errors.New("target people not found")
	}

	g.relationships[r.id] = r
	return nil
}

func (g *Organisation) AddInteraction(i Interaction) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, exists := g.people[i.peopleID]; !exists {
		return errors.New("people not found")
	}

	g.interactions[i.id] = i
	return nil
}

func (g *Organisation) FindByName(name string) (People, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, p := range g.people {
		if p.name == name {
			return p, nil
		}
	}
	return People{}, errors.New("people not found")
}

func (g *Organisation) GetPeopleRelationships(peopleID uuid.UUID) []Relationship {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	var relationships []Relationship
	for _, r := range g.relationships {
		if r.sourcePeopleID == peopleID || r.targetPeopleID == peopleID {
			relationships = append(relationships, r)
		}
	}
	return relationships
}

func (g *Organisation) GetPeopleInteractions(peopleID uuid.UUID) []Interaction {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	var interactions []Interaction
	for _, i := range g.interactions {
		if i.peopleID == peopleID {
			interactions = append(interactions, i)
		}
	}
	return interactions
}
