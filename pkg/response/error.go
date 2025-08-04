package response

type ApiErrorResponse struct {
	StatusCode int    `json:"status_code"` // HTTP status code
	Error      string `json:"error"`       // Error message
	Details    string `json:"details"`     // Additional details (optional)
}

func NewApiErrorResponse(status_code int, err string, details string) ApiErrorResponse {
	return ApiErrorResponse{
		StatusCode: status_code,
		Error:      err,
		Details:    details,
	}
}
