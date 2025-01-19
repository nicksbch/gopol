package main

type GraphData struct {
	People        []PeopleNode
	Relationships []RelationshipEdge
}

type PeopleNode struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Team      string `json:"team"`
	Level     string `json:"level"`
	Influence string `json:"influence"`
}

type RelationshipEdge struct {
	ID         string `json:"id"`
	Source     string `json:"source"`
	Target     string `json:"target"`
	Strength   string `json:"strength"`
	Type       string `json:"type"`
	Dependency string `json:"dependency"`
}

func getGraphData(org *Organisation) GraphData {
	var data GraphData

	for _, p := range org.people {
		node := PeopleNode{
			ID:        p.id.String(),
			Name:      p.name,
			Role:      p.role,
			Team:      p.currentTeam,
			Level:     p.level.String(),
			Influence: p.influence.String(),
		}
		data.People = append(data.People, node)
	}

	for _, r := range org.relationships {
		edge := RelationshipEdge{
			ID:         r.id.String(),
			Source:     r.sourcePeopleID.String(),
			Target:     r.targetPeopleID.String(),
			Strength:   r.relationshipStrength.String(),
			Type:       r.relationshipType.String(),
			Dependency: r.dependencyLevel.String(),
		}
		data.Relationships = append(data.Relationships, edge)
	}

	return data
}
