package router

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/jeffbmartinez/todo-webserver/handler"
)

const webFileDir = "./web"

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/tasks", handler.Tasks)
	// api.HandleFunc("/tasks/new", handler.NewTask)
	// api.HandleFunc("/tasks/{id}", handler.Task)

	fileServer := http.FileServer(http.Dir(webFileDir))
	router.Handle("/{pathname:.*}", fileServer)

	return router
}
