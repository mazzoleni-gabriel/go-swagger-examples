package get

import (
	"net/http"
)

// swagger:model TaskGetResponse
type TaskGetResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"value"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// swagger:parameters get-task
type _ struct {
	// in:path
	ID int `json:"id"`
}

// swagger:route GET /task/{id} Task get-task
//
// responses:
//   200: TaskGetResponse
//   404: ErrorResponse
//   500: ErrorResponse
func GetTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
