package helper

import "github.com/labstack/echo"

// EchoResp type
type EchoResp struct {
	// Error code explanations
	// 0000 = Success
	// 0001 = Error Validation
	// 0002 = Function Error
	// 0003 = Permission denied
	Code    string      `json:"code" example:"0001"`
	Message string      `json:"message" example:"this is example message"`
	Details interface{} `json:"details"`
}

// ReturnJSONresp returns response with format
func ReturnJSONresp(c echo.Context, httpcode int, code string, message string, details interface{}) error {
	if len(code) <= 0 {
		code = "0000"
	}
	x := EchoResp{
		Code:    code,
		Message: message,
		Details: details,
	}

	return c.JSON(httpcode, x)
}
