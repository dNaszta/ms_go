package main

import (
	"github.com/dNaszta/ms_go/Ch02_Lesson05_Test_Handlers/handlers"
	"net/http"
)

func main()  {
	http.HandleFunc("/example", handlers.MyHandler)
	http.ListenAndServe(":8080", nil)
}
