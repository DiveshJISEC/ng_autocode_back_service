package common

type HTTPResponse struct {
	Status       string `json:"status"`
	Data         any    `json:"data,omitempty"`
	Error        any    `json:"error,omitempty"`
	ErrorMessage string `json:"FML_ERROR_MSG,omitempty"`
}

func SuccesssResponse(data any) HTTPResponse {
	return HTTPResponse{
		Status: "success",
		Data:   data,
	}
}

func FailureResponse(error Error) HTTPResponse {
	return HTTPResponse{
		Status:       "failure",
		Error:        &error,
		ErrorMessage: error.Description,
	}
}
