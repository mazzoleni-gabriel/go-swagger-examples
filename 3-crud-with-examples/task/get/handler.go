package get

import (
	"net/http"
)

// swagger:model TaskGetResponse
type TaskGetResponse struct {
	// Task unique identifier
	// example: 123
	ID int `json:"id"`
	// Task description
	// example: implement a task API
	Description string `json:"description"`
	// Task status: TODO, WIP or DONE
	// example: TODO
	Status string `json:"value"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	// Error identifier
	// example: null_pointer_exception
	Error string `json:"error"`
	// message describing the error
	// example: Exception in thread "main" java.lang.NullPointerException
	Message string `json:"message"`
}

// swagger:parameters get-task
type _ struct {
	// Task unique identifier
	// in:path
	ID int `json:"id"`
}

// swagger:route GET /task/{id} Task get-task
//
// This is the summary
//
// This is the description, it can be longer than the summary
//
// responses:
//   200: TaskGetResponse
//   404: ErrorResponse
//   500: ErrorResponse
func GetTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
