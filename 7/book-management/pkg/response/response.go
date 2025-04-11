package response

// Response represents the standard API response format
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success returns a success response
func Success(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// Error returns an error response
func Error(message string, err error) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
	}
}
