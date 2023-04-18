package terrors

import (
	"bank_api/pkg/xruntime"
)

type tError struct {
	internalError   error
	internalMessage string
	stackTrace      *xruntime.StackTrace
	externalMessage *struct {
		Description  string `json:"description"`
		InternalCode uint16 `json:"internal_code"`
	}
}

func Raise(err error) *tError {

	tErr := &tError{
		internalError:   err,
		stackTrace:      xruntime.NewFrame(2),
		internalMessage: err.Error(),
	}

	return tErr
}

func setExternalMessage() {

}

func (e *tError) Error() string {
	return ""
}

func (e *tError) LoggerMessage() string {
	return ""
}
