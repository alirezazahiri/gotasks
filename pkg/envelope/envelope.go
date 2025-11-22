package envelope

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Response represents the standard envelope structure for all HTTP responses
type Response struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data,omitempty"`
	Error    *Error      `json:"error,omitempty"`
	Metadata *Metadata   `json:"metadata,omitempty"`
}

// Error represents error information in the response
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Metadata contains additional information about the response
type Metadata struct {
	Timestamp  string      `json:"timestamp"`
	Pagination *Pagination `json:"pagination,omitempty"`
	RequestID  string      `json:"request_id,omitempty"`
}

// Pagination contains pagination information
type Pagination struct {
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

// Success creates a successful response envelope
func Success(c *gin.Context, statusCode int, data interface{}) {
	response := Response{
		Success: true,
		Data:    data,
		Metadata: &Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
	c.JSON(statusCode, response)
}

// SuccessWithMetadata creates a successful response with custom metadata
func SuccessWithMetadata(c *gin.Context, statusCode int, data interface{}, metadata *Metadata) {
	if metadata.Timestamp == "" {
		metadata.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}
	response := Response{
		Success:  true,
		Data:     data,
		Metadata: metadata,
	}
	c.JSON(statusCode, response)
}

// SuccessWithPagination creates a successful response with pagination metadata
func SuccessWithPagination(c *gin.Context, statusCode int, data interface{}, pagination *Pagination) {
	response := Response{
		Success: true,
		Data:    data,
		Metadata: &Metadata{
			Timestamp:  time.Now().UTC().Format(time.RFC3339),
			Pagination: pagination,
		},
	}
	c.JSON(statusCode, response)
}

// ErrorResponse creates an error response envelope
func ErrorResponse(c *gin.Context, statusCode int, code, message string) {
	response := Response{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
		},
		Metadata: &Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
	c.JSON(statusCode, response)
}

// ErrorWithDetails creates an error response with additional details
func ErrorWithDetails(c *gin.Context, statusCode int, code, message, details string) {
	response := Response{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
			Details: details,
		},
		Metadata: &Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
	c.JSON(statusCode, response)
}

// BadRequest creates a 400 Bad Request response
func BadRequest(c *gin.Context, message string) {
	ErrorResponse(c, 400, "BAD_REQUEST", message)
}

// Unauthorized creates a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, 401, "UNAUTHORIZED", message)
}

// Forbidden creates a 403 Forbidden response
func Forbidden(c *gin.Context, message string) {
	ErrorResponse(c, 403, "FORBIDDEN", message)
}

// NotFound creates a 404 Not Found response
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, 404, "NOT_FOUND", message)
}

// InternalServerError creates a 500 Internal Server Error response
func InternalServerError(c *gin.Context, message string) {
	ErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", message)
}

// ValidationError creates a 422 Unprocessable Entity response
func ValidationError(c *gin.Context, message, details string) {
	ErrorWithDetails(c, 422, "VALIDATION_ERROR", message, details)
}

// Created creates a 201 Created response
func Created(c *gin.Context, data interface{}) {
	Success(c, 201, data)
}

// OK creates a 200 OK response
func OK(c *gin.Context, data interface{}) {
	Success(c, 200, data)
}

// NoContent creates a 204 No Content response
func NoContent(c *gin.Context) {
	response := Response{
		Success: true,
		Metadata: &Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
	c.JSON(204, response)
}
