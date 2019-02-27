package main

import (
	"log"
	"html/template"
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
	homeTemp := template.Must(template.ParseFiles("html/templates/layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	homeTemp.Execute(writer, data)
}
func generate(writer http.ResponseWriter, request *http.Request) {
	generateTemp := template.Must(template.ParseFiles("html/generate.html"))

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



	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8880", nil)
}