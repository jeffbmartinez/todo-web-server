package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffbmartinez/log"
)

// Task handles requests to the /tasks/{id} endpoint.
func TaskHandler(response http.ResponseWriter, request *http.Request) {
	handler := BasicResponse(http.StatusMethodNotAllowed)

	switch request.Method {
	case "GET":
		handler = showTask
	}

	handler(response, request)
}

func showTask(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	taskID := vars["id"]

	endpoint := Services.Storage.Endpoint + "/tasks/" + taskID

	storageResponse, err := http.Get(endpoint)
	if err != nil {
		log.Errorf("No response from endpoint '%v' (%v)", endpoint, err)
		WriteBasicResponse(http.StatusInternalServerError, response)
		return
	}

	var task Task
	defer storageResponse.Body.Close()
	decoder := json.NewDecoder(storageResponse.Body)
	err = decoder.Decode(&task)
	if err != nil {
		log.Errorf("Couldn't read body as json in response from '%v' (%v)", endpoint, err)
		WriteBasicResponse(http.StatusInternalServerError, response)
		return
	}

	taskTemplate, err := template.ParseFiles("./templates/task.html", "./templates/newtask.html")
	if err != nil {
		log.Errorf("Couldn't read task template")
		WriteBasicResponse(http.StatusInternalServerError, response)
		return
	}

	fmt.Println(task)

	taskTemplate.ExecuteTemplate(response, "task", task)
}
