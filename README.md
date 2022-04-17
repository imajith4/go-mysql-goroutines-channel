# go-mysql-goroutines-channel
Sample code snippet to familiarize golang . Concept of goroutines , channels, concurrency , interface, slice, error handling, Blank Identifier etc is implemented in a sample scenario

# Preface
This repository is the sample of web application using golang. This code snippet used mysql as DBMS , GORM  as ORM and .tmpl extension in frontend (go template file)

# Install
Perform the following steps:
  1. Download and install Golang - https://go.dev/ .
  2.  go get gorm.io/driver/mysql 
  3.  go get -u gorm.io/gorm

  NB: Please refer https://gorm.io/ before installing the ORM packages
  
 # WorkFlow
 - This application has a basic interface Event which has 2 abstract functions All(), New()
 - 2 structures defined (Encounter, AuditLog) implements the interface Event
 - There is a form to add bulk data 
 - in main package there is 2  goroutines defined 
        *  poolDataToChannel -  To pool the bulk data into a channel
        *  insertDataFromPool -  To insert data into encounter , audit log from the channel   

 - When the form is filled and submitted 2 goroutines run concurrenlty ie.. fetching and insertion is done concorrently

# Run the application
_> cd gosampleportal

gosampleportal _> run go .  (in windows)
gosampleportal _> run go *.go  (in linux)


# Sample console output
D:\code\gosampleportal>go run .
Userid -  1 , Location -  john location   added to channel
Userid -  1 , Location -  john location  . Inserted to encounter
Userid -  1 , Location -  john location   Logtype - Portal , Module - Calendar. Inserted to Audit Log
Userid -  2 , Location -  joe location  added to channel
Userid -  2 , Location -  joe location . Inserted to encounter
Userid -  2 , Location -  joe location  Logtype - Portal , Module - Calendar. Inserted to Audit Log
Userid -  3 , Location -  jane location  added to channel
Userid -  3 , Location -  jane location . Inserted to encounter
Userid -  3 , Location -  jane location  Logtype - Portal , Module - Calendar. Inserted to Audit Log


#concurrency vs parallelism
If there’s one thing most people know about Go, is that it is designed for concurrency. No introduction to Go is complete without a demonstration of its goroutines and channels.

But when people hear the word concurrency they often think of parallelism, a related but quite distinct concept. In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

https://go.dev/blog/waza-talk




 
 
