package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jeffbmartinez/todo-webserver/handler"
)

const webFileDir = "./web"

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/tasks", handler.ApiTasks)
	// api.HandleFunc("/tasks/new", handler.ApiNewTask)
	// api.HandleFunc("/tasks/{id}", handler.ApiTask)

	router.HandleFunc("/tasks", handler.Tasks)

	fileServer := http.FileServer(http.Dir(webFileDir))
	router.Handle("/{pathname:.*}", fileServer)

	return router
}
