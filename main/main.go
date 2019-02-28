package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"../tools"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}



func logger (f http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer,request)
	}
}


func home(writer http.ResponseWriter, request *http.Request) {

	templateBox, err := rice.FindBox("../html")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("home.html")
	if err != nil {
		log.Fatal(err)
	}
	// parse and execute the template
	layoutTemp, err := template.New("home").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	data , _ := tools.IOReadDir("./")
	layoutTemp.Execute(writer, data)
}


func generate(writer http.ResponseWriter, request *http.Request) {

	templateBox, err := rice.FindBox("../html")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("generate.html")
	if err != nil {
		log.Fatal(err)
	}
	// parse and execute the template
	generateTemp, err := template.New("generate").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}


	if request.Method != http.MethodPost {
		generateTemp.Execute(writer, nil)
		return
	}

	details := ContactDetails{
		Email:   request.FormValue("email"),
		Subject: request.FormValue("subject"),
		Message: request.FormValue("message"),
	}

	// do something with details
	_ = details

	generateTemp.Execute(writer, struct{ Success bool }{true})
}



func main() {

	http.HandleFunc("/", logger(home))

	http.HandleFunc("/generate", logger(generate))

	box := rice.MustFindBox("../assets")
	cssFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	http.Handle("/static/", cssFileServer)


	project := http.StripPrefix("/load/", http.FileServer(rice.MustFindBox("../").HTTPBox()))
	http.Handle("/load/", project)

	fmt.Println(tools.IOReadDir("./"))

	http.ListenAndServe(":8880", nil)

}