package delete

import "net/http"

// swagger:model TaskDeleteResponse
type TaskDeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// swagger:parameters delete-task
type _ struct {
	// in:path
	ID int `json:"id"`
}

// swagger:route DELETE /task/{id} Task delete-task
//
// responses:
//   204: TaskDeleteResponse
//   404: ErrorResponse
//   500: ErrorResponse
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
