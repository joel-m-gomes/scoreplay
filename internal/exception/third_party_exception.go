package exception

import (
	"fmt"
)

type ThirdPartyException struct {
	Source     string
	StatusCode int
	Message    string
}

func (e *ThirdPartyException) Error() string {
	return fmt.Sprintf("Third party error: source %s, status code %d, message: %s", e.Source, e.StatusCode, e.Message)
}

func NewThirdPartyException(source string, statusCode int, message string) *ThirdPartyException {
	return &ThirdPartyException{
		Source:     source,
		StatusCode: statusCode,
		Message:    message,
	}
}
