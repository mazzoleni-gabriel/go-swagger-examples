package list

import "net/http"

// swagger:model TaskListResponse
type TaskListResponse struct {
	List []Task `json:"list"`
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"value"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// swagger:parameters list-task
type _ struct {
	// in:query
	Status string `json:"status"`
	// in:header
	XRequestId string `json:"x-request-id"`
}

// swagger:route GET /task/list Task list-task
//
// responses:
//   200: TaskListResponse
//   404: ErrorResponse
//   500: ErrorResponse
func ListTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
