package main

import (
	"gosampleportal/model"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./templates/home.tmpl")
	parsedTemplate.Execute(w, nil)

}

//------------encounter related functions---------
func AddEncounter(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./templates/add.tmpl")
	parsedTemplate.Execute(w, nil)

}

func GetAllEncounters(w http.ResponseWriter, r *http.Request) {
	enc := model.Encounter{}
	getall(enc, w, r)
}

func getall(g model.Encounter, w http.ResponseWriter, r *http.Request) {
	g.All(w, r)
	//fmt.Println(g.perim())
}

func SaveEncounter(w http.ResponseWriter, r *http.Request) {
	enc := model.Encounter{}
	saveEnc(enc, w, r)
}
func saveEnc(g model.Encounter, w http.ResponseWriter, r *http.Request) {
	g.New(w, r)
	//fmt.Println(g.perim())
}

//------------encounter related functions---------

//------------log related functions---------
func GetAllLog(w http.ResponseWriter, r *http.Request) {
	log := model.AuditLog{}
	getallLog(log, w, r)
}

func getallLog(g model.AuditLog, w http.ResponseWriter, r *http.Request) {
	g.All(w, r)
	//fmt.Println(g.perim())
}

func SaveLog(w http.ResponseWriter, r *http.Request) {
	log := model.AuditLog{}
	saveLog(log, w, r)
}
func saveLog(g model.AuditLog, w http.ResponseWriter, r *http.Request) {
	g.New(w, r)
	//fmt.Println(g.perim())
}

//------------log related functions---------
