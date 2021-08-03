package common

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	// common
	ErrNoData   = errors.New("No Data")
	ErrNotFound = errors.New("Not Found")

	// Account
	ErrIdInvalid = errors.New("id is empty")
)

// Error Message
func ErrorMsg(c echo.Context, status int, err error) {
	errMsg := messageFormat{
		StatusCode: status,
		Message:    err.Error(),
	}
	c.JSON(status, echo.Map{"error": errMsg})
}

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
}
