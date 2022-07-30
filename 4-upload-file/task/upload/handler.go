package upload

import "net/http"

// swagger:model TaskUploadResponse
type TaskUploadResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// swagger:parameters upload-task
type _ struct {
	// in:formData
	// swagger:file
	File interface{} `json:"file"`
	// in:header
	Boundary string `json:"boundary"`
}

// swagger:route POST /task/upload Task upload-task
//
// Consumes:
// - multipart/form-data
//
// responses:
//   201: TaskUploadResponse
//   404: ErrorResponse
//   500: ErrorResponse
func UploadTask(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
