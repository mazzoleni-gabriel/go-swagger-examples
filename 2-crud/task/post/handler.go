package post

import (
	"net/http"
)

// swagger:model TaskPostRequest
type TaskPostRequest struct {
	Description string `json:"description"`
}

// swagger:model TaskPostResponse
type TaskPostResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// swagger:parameters post-task
type _ struct {
	// in:body
	// required: true
	Body TaskPostRequest
}

// swagger:route POST /task/ Task post-task
//
// responses:
//   201: TaskPostResponse
//   404: ErrorResponse
//   500: ErrorResponse
func PostTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
