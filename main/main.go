package main

import (
	"../gegosc"
	_ "../routers"
	"net/http"
)




func main() {

	go http.ListenAndServe(":8880", nil)
	 gegosc.ProxyStart(8888,8880)




}