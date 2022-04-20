package utils

import (
	"encoding/json"
	"net/http"
)



type ErrorHandlerInterface interface {
	MessageErr() string
	Status() int
	Error() string
}
type ErrorHandlerBuilder interface {
	WithMessageErr(str string) ErrorHandlerBuilder
	WithStatus(i int) ErrorHandlerBuilder
	WithError(e string) ErrorHandlerBuilder
	Builder() messageErr
}
type messageErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}


func (e *messageErr) MessageErr() string {
	return e.ErrMessage
}

func (e *messageErr) Status() int {
	return e.ErrStatus
}

func (e *messageErr) Error() string {
	return e.ErrError
}

func (e *messageErr) WithMessageErr(str string) ErrorHandlerBuilder {
	e.ErrMessage = str
	return e
}
// Builder implements ErrorHandlerBuilder
func (mess messageErr) Builder() messageErr {
	return mess
}

// WithError implements ErrorHandlerBuilder
func (mess *messageErr) WithError(e string) ErrorHandlerBuilder {
	mess.ErrError = e
	return mess
}

// WithStatus implements ErrorHandlerBuilder
func (mess *messageErr) WithStatus(i int) ErrorHandlerBuilder {
	mess.ErrStatus = int(i)
	return mess
}
func NewNotFoundErr(str string) ErrorHandlerInterface {
	return &messageErr{ErrMessage: str, ErrStatus: http.StatusNotFound, ErrError: str}
}
func NewBadRequestError(message string) ErrorHandlerInterface {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}
func NewUnprocessibleEntityError(message string) ErrorHandlerInterface {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}

func NewApiErrFromBytes(body []byte) (ErrorHandlerInterface, error) {
	var result messageErr
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func NewInternalServerError(message string) ErrorHandlerInterface {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "server_error",
	}
}
