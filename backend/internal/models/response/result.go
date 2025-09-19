package response

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewErrorMessage(message string) *Result {
	return &Result{
		Code:    1,
		Message: message,
	}
}

func NewErrorData(message string, data any) *Result {
	return &Result{
		Code:    1,
		Message: message,
		Data:    data,
	}
}

func NewSuccessMessage(message string) *Result {
	return &Result{
		Code:    0,
		Message: message,
	}
}

func NewSuccessData(message string, data any) *Result {
	return &Result{
		Code:    0,
		Message: message,
		Data:    data,
	}
}
