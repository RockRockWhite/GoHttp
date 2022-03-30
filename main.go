package main

import "net/http"

type myHandler struct {
}

func (m myHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello world"))
	//})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &myHandler{},
	}
	server.ListenAndServe()
	//http.ListenAndServe("localhost:8080", nil)
}
