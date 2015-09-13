package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jeffbmartinez/log"
)

// Tasks handles requests to the /tasks endpoint.
func Tasks(response http.ResponseWriter, request *http.Request) {
	handler := BasicResponse(http.StatusMethodNotAllowed)

	switch request.Method {
	case "GET":
		handler = showTasks
	}

	handler(response, request)
}

func showTasks(response http.ResponseWriter, request *http.Request) {
	endpoint := Services.Storage.Endpoint + "/tasks"

	storageResponse, err := http.Get(endpoint)
	if err != nil {
		log.Errorf("No response from endpoint '%v' (%v)", endpoint, err)
		WriteBasicResponse(http.StatusInternalServerError, response)
		return
	}

	var tasks []Task
	defer storageResponse.Body.Close()
	decoder := json.NewDecoder(storageResponse.Body)
	err = decoder.Decode(&tasks)
	if err != nil {
		log.Errorf("Couldn't read body as json in response from '%v' (%v)", endpoint, err)
		WriteBasicResponse(http.StatusInternalServerError, response)
		return
	}

	tasksTemplate, err := template.ParseFiles("./templates/tasks.html")
	if err != nil {
		log.Errorf("Couldn't read tasks template")
		WriteBasicResponse(http.StatusInternalServerError, response)
		return
	}

	fmt.Println(tasks)

	tasksTemplate.Execute(response, tasks)
}
