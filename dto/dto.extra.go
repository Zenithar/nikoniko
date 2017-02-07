package dto

import (
	"fmt"
)

// Error method makes the Error type implement the error interface
func (e *Error) Error() (message string) {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}
