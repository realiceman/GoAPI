package main

import (
	"GoAPI/src/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}

}
