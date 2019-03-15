package main

import (
	_ "../routers"
	"net/http"
)







func main() {



	http.ListenAndServe(":8880", nil)

}