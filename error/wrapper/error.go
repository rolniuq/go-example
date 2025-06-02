package wrapper

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code    int
	Type    string
	Message string
	cause   error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	var js json.RawMessage
	if json.Unmarshal([]byte(e.Message), &js) == nil {
		return string(js)
	}

	if e.cause != nil {
		return fmt.Sprintf("%s: %s (cause: %s)", e.Type, e.Message, e.cause)
	}

	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *Error) Is(target error) bool {
	if e == nil {
		return false
	}

	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return e.Type == t.Type
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.cause
}

func NewError(code int, typ string, message string, cause error) *Error {
	return &Error{
		Code:    code,
		Type:    typ,
		Message: message,
		cause:   cause,
	}
}
