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

// func NewInternalError(details string) *AppError {
// 	return &AppError{
// 		Code:    http.StatusInternalServerError,
// 		Message: "Internal server error",
// 		Details: details,
// 	}
// }

func NewNotFoundError(details string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
		Details: details,
	}
}

// func NewBadRequestError(details string) *AppError {
// 	return &AppError{
// 		Code:    http.StatusBadRequest,
// 		Message: "Bad request",
// 		Details: details,
// 	}
// }

func NewUnauthorizedError(details string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Details: details,
	}
}

func NewForbiddenError(details string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
		Details: details,
	}
}
