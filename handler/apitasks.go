package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeffbmartinez/log"
)

/*
Task represents a task or a subtask. Tasks can be stand-alone todo items
or they can be broken down into subtasks, which can then have their own
subtasks, and so on. Subtasks have parents which they can be grouped into.
Tasks can have multiple parents to cover the possibility of a single task
accomplishing two parent tasks, for example "clean room" can be a subtask
for "clean house" as well as "prepare for parents' visit".

A task with no parents is called a "root" task.
*/
type Task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:complete`

	CreatedDate  int64 `json:"createdDate"`
	ModifiedDate int64 `json:"modifiedDate"`
	DueDate      int64 `json:"dueDate"`

	Categories []string `json:"categories"`

	Subtasks []*Task `json:"subtasks"`
}

// ApiTasks handles requests to the /tasks/{id} endpoint.
func ApiTasks(response http.ResponseWriter, request *http.Request) {
	handler := BasicResponse(http.StatusMethodNotAllowed)

	switch request.Method {
	case "GET":
		handler = getTasks
	}

	handler(response, request)
}

func getTasks(response http.ResponseWriter, request *http.Request) {
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

	WriteJSONResponse(response, tasks, http.StatusOK)
}
