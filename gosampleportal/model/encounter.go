package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
)

// Encounter model
type Encounter struct {
	gorm.Model
	ID     int    `gorm:"type:int"`
	Name   string `gorm:"type:varchar(250)"`
	Userid int64  `gorm:"type:int"`
}

// Implementing methods of  the event interface
func (e Encounter) All(w http.ResponseWriter, r *http.Request) {
	dsn := "root:@tcp(127.0.0.1:3306)/gosample?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	encounter := []Encounter{}
	db.Find(&encounter)
	parsedTemplate, _ := template.ParseFiles("./templates/encounter.tmpl")
	if err := parsedTemplate.Execute(w, encounter); err != nil {
		panic(err)
	}
}
func (e Encounter) New(w http.ResponseWriter, r *http.Request) {
	var c chan Event = make(chan Event)

	dsn := "root:@tcp(127.0.0.1:3306)/gosample?parseTime=true"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	bulkSlice := formToSlice(r)
	go poolDataToChannel(bulkSlice, c)
	go insertDataFromPool(c)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func poolDataToChannel(eventslice []Event, c chan Event) {
	for _, e := range eventslice {
		c <- e
		fmt.Println("Userid - ", e.Userid, ", Location - ", e.Name, " added to channel")
	}
}

func insertDataFromPool(c chan Event) {
	for {
		element := <-c

		dsn := "root:@tcp(127.0.0.1:3306)/gosample?parseTime=true"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		encounter := Encounter{Name: element.Name, Userid: element.Userid}
		err1 := db.Create(&encounter)
		if err1 != nil {
		}
		fmt.Println("Userid - ", element.Userid, ", Location - ", element.Name, ". Inserted to encounter")

		auditLog := AuditLog{Name: element.Name, Userid: element.Userid, Logtype: "Portal", Module: "Calendar"}
		db.Create(&auditLog)
		if err != nil {
			panic(err)
		}
		fmt.Println("Userid - ", element.Userid, ", Location - ", element.Name, " Logtype - Portal , Module - Calendar. Inserted to Audit Log")

		//time.Sleep(time.Second * 1)
	}

}

func formToSlice(r *http.Request) []Event {
	//myslice := Event{}[1:3]

	//var my_slice_1 = []string{"Geeks", "for", "Geeks"}
	location1 := r.FormValue("location1")
	userid1, _ := strconv.ParseInt(r.FormValue("userid1"), 10, 64)

	enc1 := Event{1, location1, userid1}
	//go addToChannel(c, enc1, w)
	location2 := r.FormValue("location2")
	userid2, _ := strconv.ParseInt(r.FormValue("userid2"), 10, 64)

	enc2 := Event{2, location2, userid2}
	//go addToChannel(c, enc2, w)
	location3 := r.FormValue("location3")
	userid3, _ := strconv.ParseInt(r.FormValue("userid3"), 10, 64)

	enc3 := Event{3, location3, userid3}

	myslice := []Event{enc1, enc2, enc3}
	return myslice

}
