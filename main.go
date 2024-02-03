// cmd/wizard-web-app/main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Wizard struct {
	Name  string `json:"name"`
	House string `json:"house"`
	Age   int    `json:"age"`
}

var wizards = []Wizard{
	{Name: "Harry Potter", House: "Gryffindor", Age: 20},
	{Name: "Hermione Granger", House: "Gryffindor", Age: 21},
	{Name: "Ron Weasley", House: "Gryffindor", Age: 22},
	{Name: "Draco Malfoy", House: "Slytherin", Age: 23},
}

func getWizardsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wizards)
}

func getWizardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	name := params["name"]

	for _, wizard := range wizards {
		if wizard.Name == name {
			json.NewEncoder(w).Encode(wizard)
			return
		}
	}

	http.Error(w, "Wizard not found", http.StatusNotFound)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Wizard Web App is healthy. Author: ChatGPT")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/wizards", getWizardsHandler).Methods("GET")
	router.HandleFunc("/wizards/{name}", getWizardHandler).Methods("GET")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

/*http://localhost:8080/wizards
http://localhost:8080/wizards/Harry%20Potter*/
