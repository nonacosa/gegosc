package controllers
import (
	"../models"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"net/http"
)

func Generate(writer http.ResponseWriter, request *http.Request) {

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

	details := models.ContactDetails{
		Email:   request.FormValue("email"),
		Subject: request.FormValue("subject"),
		Message: request.FormValue("message"),
	}

	// do something with details
	_ = details

	generateTemp.Execute(writer, struct{ Success bool }{true})
}


