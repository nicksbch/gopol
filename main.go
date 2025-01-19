package main

import (
	"fmt"
)

func loadOrganisation() *Organisation {
	org := NewOrganisation()
	org.AddPeople(NewPeople("John Doe", "Manager", "Team A", SeniorEmployeeLevel, HighInfluence, "A very influential person"))
	org.AddPeople(NewPeople("Jane Doe", "Manager", "Team B", SeniorEmployeeLevel, HighInfluence, "A very influential person"))

	people1, _ := org.FindByName("John Doe")
	people2, _ := org.FindByName("Jane Doe")
	org.AddRelationship(NewRelationship(people1, people2, StrongRelationshipStrength, Manager, LowDependencyLevel, "A very strong relationship"))

	return org
}

func main() {
	org := loadOrganisation()
	people1, _ := org.FindByName("John Doe")

	rsaPubKey, err := loadRSAPublicKey("test-pk-rs256.rsa.pub")
	fmt.Print(people1.Encrypt(rsaPubKey))
}
