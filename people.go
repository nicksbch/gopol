package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type People struct {
	id              uuid.UUID
	name            string
	role            string
	currentTeam     string
	level           Level
	influence       Influence
	notes           string
	attributes      []string
	lastInteraction *time.Time
}

func NewPeople(name, role, team string, level Level, influence Influence, notes string) People {
	return People{
		id:              uuid.New(),
		name:            name,
		role:            role,
		currentTeam:     team,
		level:           level,
		influence:       influence,
		notes:           notes,
		attributes:      []string{},
		lastInteraction: nil,
	}
}

func (p People) String() string {
	return fmt.Sprintf("%s (%s) %s %s, %s", p.name, p.level, p.role, p.currentTeam, p.influence)
}

func (p People) Encrypt(rsaPublicKey *rsa.PublicKey) ([]byte, error) {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, rsaPublicKey, []byte(p.String()), nil)
	return ciphertext, err
}
