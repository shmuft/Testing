package main

import (
	"net/http"
	"html/template"
	"time"
	"log"
)

type Application struct {
	dbMongo *dbMongo
}
var app Application
type user struct {
	UserName string
	Ank      anketa
	T1       []testAnswer
	T2       []testAnswer
	T3       []testAnswer
	Level    int
	Password []byte
}

type testAnswer struct {
	Id      int
	Answers []int
}

type anketa struct {
	Name            string
	Sex             string
	Age             int
	Education       int
	Experience      int
	HowManyPeople   int
	HowLongWorkProf int
	HowLongWork     int
}

type test2 struct {
	title string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 1 * 60 * 60 * 3

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
	InitializeDBTest1()
}

func main() {
	dbMongo, err := NewMongoDb()
	if err != nil {
		log.Println(err)
		return
	}
	app = Application{dbMongo: dbMongo}
	defer app.dbMongo.CloseMongoDb()

	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/about", AboutHandle)
	http.HandleFunc("/anketa", AnketaHandle)
	http.HandleFunc("/test1", Test1Handle)
	http.HandleFunc("/test2", Test2Handle)
	http.HandleFunc("/test3", Test3Handle)
	http.HandleFunc("/signup", SignupHandle)
	http.HandleFunc("/login", LoginHandle)
	http.HandleFunc("/logout", LogoutHandle)
	http.HandleFunc("/get_data_test1", GetDataTest1)
	http.HandleFunc("/get_data_test2", GetDataTest2)
	http.HandleFunc("/get_data_test3", GetDataTest3)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println(http.ListenAndServe(":8080", nil))
}
