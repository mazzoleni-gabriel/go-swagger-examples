package put

import (
	"net/http"
)

// swagger:model TaskPutRequest
type TaskPutRequest struct {
	Description string `json:"description"`
	Status      string `json:"value"`
}

// swagger:model TaskPutResponse
type TaskPutResponse struct {
	Status int `json:"status"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// swagger:parameters Put-task
type _ struct {
	// in:path
	ID int `json:"id"`

	// in:body
	// required: true
	Body TaskPutRequest
}

// swagger:route Put /task/{id} Task Put-task
//
// responses:
//   201: TaskPutResponse
//   404: ErrorResponse
//   500: ErrorResponse
func PutTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
