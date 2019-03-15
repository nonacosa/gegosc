package controllers

import (
	"../tools"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {

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
