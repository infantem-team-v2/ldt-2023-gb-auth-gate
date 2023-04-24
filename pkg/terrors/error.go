package terrors

import (
	"bank_api/pkg/xruntime"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

// tError is an extended type of error which has stacktrace and mechanisms for better debugging and handling errors
type tError struct {
	// internal fields for our purposes
	internalError error
	stackTrace    *xruntime.StackTrace

	// external fields to send to client
	statusCode      int
	externalMessage *externalMessage
}

// Init cache with errors' external messages to have fast way to get information about them
func Init() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fErr, err := os.Open(fmt.Sprintf("%s/test/errors.yaml", workDir))
	if err != nil {
		panic(err)
	}
	bErr, err := io.ReadAll(fErr)
	if err != nil {
		panic(err)
	}
	if len(bErr) == 0 {
		panic(errors.New("no information in errors.yaml file"))
	}

	if err := yaml.Unmarshal(bErr, &externalMessagesMap); err != nil {
		panic(err)
	}

}

func getExternalMessage(internalCode uint32) (extMsg *externalMessage, code int, err error) {
	serExtMsg := externalMessagesMap[internalCode]
	if serExtMsg == nil {
		return &defaultErrorMessage, defaultStatusCode, errors.New("there's no cached errors in system")
	}
	return &externalMessage{
		Description:  serExtMsg.Message,
		InternalCode: internalCode,
	}, serExtMsg.StatusCode, nil
}

func Raise(err error, internalCode uint32) IError {
	extMsg, statusCode, iErr := getExternalMessage(internalCode)
	if err == nil {
		err = errors.New(extMsg.Description)
	}
	if iErr != nil {
		err = errors.Wrap(iErr, err.Error())
	}
	tErr := &tError{
		internalError: err,
		stackTrace:    xruntime.NewFrame(2),

		statusCode:      statusCode,
		externalMessage: extMsg,
	}

	return tErr
}

func (e *tError) Wrap(err IError) IError {
	if tErr, ok := err.(*tError); ok {
		e.internalError = errors.Wrap(tErr.internalError, fmt.Sprintf("%s", e.internalError.Error()))
	}
	return e
}

func stacktraceToString(st *xruntime.StackTrace, doNewLine bool) (serializedFrames string) {
	frames := st.Frames()
	for _, frame := range frames {
		serializedFrames += fmt.Sprintf("File: %s; Function: %s;  Line: %d;",
			frame.File, frame.Function, frame.Line)
		if doNewLine {
			serializedFrames += "\n"
		}
	}
	return serializedFrames
}

func (e *tError) Error() string {

	return ""
}

func (e *tError) LoggerMessage() string {
	serializedFrames := stacktraceToString(e.stackTrace, false)
	return fmt.Sprintf(labels,
		e.statusCode,
		e.externalMessage.InternalCode, e.externalMessage.Description,
		e.internalError.Error(),
		serializedFrames)
}
