package main

import (
	"html/template"
	"log"
	"net/http"
)

func loadOrganisation() *Organisation {
	org := NewOrganisation()
	org.AddPeople(NewPeople("John Doe", "Manager", "Team A", SeniorEmployeeLevel, HighInfluence, "A very influential person"))
	org.AddPeople(NewPeople("Jane Doe", "Manager", "Team B", SeniorEmployeeLevel, HighInfluence, "A very influential person"))

	people1, _ := org.FindByName("John Doe")
	people2, _ := org.FindByName("Jane Doe")
	org.AddRelationship(NewRelationship(people1, people2, StrongRelationshipStrength, Manager, LowDependencyLevel, "A very strong relationship"))
	org.AddInteraction(NewInteraction(people1, InteractionTypeMeeting, InteractionSentimentPositive, InteractionImpactPositive, "A very strong relationship"))
	return org
}

func main() {
	org := loadOrganisation()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/graph.html",
		))

		data := getGraphData(org)
		err := tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
