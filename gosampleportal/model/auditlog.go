package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

// AuditLog model
type AuditLog struct {
	gorm.Model
	ID      int    `gorm:"type:int"`
	Name    string `gorm:"type:varchar(250)"`
	Userid  int64  `gorm:"type:int"`
	Logtype string `gorm:"type:varchar(250)" `
	Module  string `gorm:"type:varchar(250)" `
}

// Implementing methods of  the event interface
func (a AuditLog) All(w http.ResponseWriter, r *http.Request) {
	dsn := "root:@tcp(127.0.0.1:3306)/gosample?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	auditLog := []AuditLog{}
	db.Find(&auditLog)
	parsedTemplate, _ := template.ParseFiles("./templates/log.tmpl")
	if err := parsedTemplate.Execute(w, auditLog); err != nil {
		panic(err)
	}
}
func (e AuditLog) New(w http.ResponseWriter, r *http.Request) {
	dsn := "root:@tcp(127.0.0.1:3306)/gosample?parseTime=true"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
