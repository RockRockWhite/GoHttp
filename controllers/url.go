package controllers

import (
	"fmt"
	"net/http"
)

func registerHomeRoutes() {
	http.HandleFunc("/url", handleHomes)
	//http.HandleFunc("/url/", handleHomes)
}

func handleHomes(writer http.ResponseWriter, request *http.Request) {
	//request.ParseMultipartForm(1024)
	//file, _, _ := request.FormFile("hello")
	//file_bytes, _ := ioutil.ReadAll(file)
	//
	//writer.Header().Set("location", "www.google.com")
	//writer.Header().Set("Content-Type", "application/json")
	//writer.WriteHeader(http.StatusBadGateway)
	//j, _ := json.Marshal(request.MultipartForm)
	//fmt.Fprintln(writer, j)
	//fmt.Fprintln(writer, "j")
	//fmt.Println(writer, string(file_bytes))
	fmt.Fprintln(writer, "url")
}