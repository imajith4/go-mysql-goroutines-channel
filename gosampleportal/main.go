package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", GetAllLog)
	http.HandleFunc("/all", GetAllEncounters)
	http.HandleFunc("/add", AddEncounter)
	http.HandleFunc("/save", SaveEncounter)

	http.ListenAndServe(":8080", nil)
}
