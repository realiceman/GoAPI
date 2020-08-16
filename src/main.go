package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}

}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Asset not found\n"))
		return
	}
  writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Running API v1\n"))
}