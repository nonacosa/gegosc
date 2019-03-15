package gegosc

import (
	"log"
	"net/http"
)

func Handler (f http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer,request)
	}
}
