package errors

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// 由 validator 处理
// 400
// func NewBadRequestError(details string) *AppError {
// 	return &AppError{
// 		Code:    http.StatusBadRequest,
// 		Message: "Bad request",
// 		Details: details,
// 	}
// }

// 401
func NewUnauthorizedError(details string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Details: details,
	}
}

// 403
func NewForbiddenError(details string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
		Details: details,
	}
}

// 404
func NewNotFoundError(details string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
		Details: details,
	}
}

// 409
func NewConflictError(details string) *AppError {
	return &AppError{
		Code:    http.StatusConflict,
		Message: "Conflict",
		Details: details,
	}
}

// 422
func NewUnprocessableEntityError(details string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
		Details: details,
	}
}

// 以上错误之外统一为500并统一处理
// 500
// func NewInternalError(details string) *AppError {
// 	return &AppError{
// 		Code:    http.StatusInternalServerError,
// 		Message: "Internal server error",
// 		Details: details,
// 	}
// }
