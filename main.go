package main

import (
	"go_web/controllers"
	"net/http"
	"time"
)

type helloHandler struct {
}

func (m helloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(10 * time.Second)
	writer.Write([]byte("hello world"))
}

type aboutHandler struct {
}

func (m aboutHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("about"))
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello world"))
	//})

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	http.ServeFile(writer, request, "wwwroot"+request.URL.Path)
	//})
	//
	//http.Handle("/hello", &helloHandler{})
	//http.Handle("/about", &aboutHandler{})
	//http.Handle("/home", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("welcome"))
	//}))
	controllers.RegisterRoutes()

	server := http.Server{Addr: ":8080"}
	server.ListenAndServe()
}
