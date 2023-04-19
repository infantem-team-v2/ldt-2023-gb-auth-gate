package terrors

import (
	"bank_api/pkg/xruntime"
)

type tError struct {
	// internal fields for our purposes
	internalError   error
	internalMessage string
	stackTrace      *xruntime.StackTrace

	// external fields to send to client
	statusCode      int
	externalMessage *struct {
		Description  string `json:"description"`
		InternalCode uint16 `json:"internal_code"`
	}
}

func Raise(err error) IError {

	tErr := &tError{
		internalError:   err,
		stackTrace:      xruntime.NewFrame(2),
		internalMessage: err.Error(),
	}

	return tErr
}

func (e *tError) Wrap(err IError) IError {

	return e
}

func setExternalMessage() {

}

func (e *tError) Error() string {
	return ""
}

func (e *tError) LoggerMessage() string {
	return ""
}
