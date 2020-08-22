package handlers

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

// users route
func UsersRouter(w http.ResponseWriter, r *http.Request)  {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			usersPostOne(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
		}
	}

	path = strings.TrimPrefix(path, "/users/")
	fmt.Println("path is ",path)
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}
	id := bson.ObjectIdHex(path)
	fmt.Println("id is ", id)
	switch r.Method {
	case http.MethodGet:
		usersGetOne(w, r, id)
		return
	case http.MethodPut:
		usersPutOne(w, r, id)
		return
	case http.MethodPatch:
		usersPatchOne(w, r, id)
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
