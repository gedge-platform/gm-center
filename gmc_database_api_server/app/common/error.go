package common

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	// common
	ErrNoData            = errors.New("No Data")
	ErrNotFound          = errors.New("Not Found")
	ErrClusterNotFound   = errors.New("Cluster Not Found")
	ErrWorkspaceNotFound = errors.New("Workspace Not Found")
	ErrWorkspaceInvalid  = errors.New("Workspace Empty")
	ErrDetailNameInvalid = errors.New("Detail Name Empty")
	ErrClusterInvalid    = errors.New("Required Cluster Name")
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
