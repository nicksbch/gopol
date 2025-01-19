package main

type RelationshipStrength int

const (
	WeakRelationshipStrength RelationshipStrength = iota
	MediumRelationshipStrength
	StrongRelationshipStrength
	UnknownRelationshipStrength
)

var RelationshipStrengthName = map[RelationshipStrength]string{
	WeakRelationshipStrength:    "weak",
	MediumRelationshipStrength:  "medium",
	StrongRelationshipStrength:  "strong",
	UnknownRelationshipStrength: "unknown",
}

func (i RelationshipStrength) String() string {
	return RelationshipStrengthName[i]
}

type RelationshipType int

const (
	Manager RelationshipType = iota
	Teammate
	Subordinate
	Colleague
)

var RelationshipTypeName = map[RelationshipType]string{
	Manager:     "manager",
	Teammate:    "teammate",
	Subordinate: "subordinate",
	Colleague:   "colleague",
}

func (i RelationshipType) String() string {
	return RelationshipTypeName[i]
}

type DependencyLevel int

const (
	LowDependencyLevel DependencyLevel = iota
	MediumDependencyLevel
	HighDependencyLevel
	UnknownDependencyLevel
)

var DependencyLevelName = map[DependencyLevel]string{
	LowDependencyLevel:     "low",
	MediumDependencyLevel:  "medium",
	HighDependencyLevel:    "high",
	UnknownDependencyLevel: "unknown",
}

func (i DependencyLevel) String() string {
	return DependencyLevelName[i]
}
