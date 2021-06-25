package niceErrors

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
)

type NiceErrors struct {
	errorMessage           string
	userFacingErrorMessage string
	stackTrace             string
	errorLevel             errLevel
	errorType              ErrorType
}

type ErrorType string

const defaultErrorType ErrorType = "defaultErrorType"

type niceErrorsJson struct {
	UserFacingErrorMessage string `json:"error_message"`
}

type errLevel int

const (
	DEBUG errLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func New(errorMessage string, userErrorMessage string, errorType ErrorType, errorLevel errLevel) *NiceErrors {
	if errorLevel < 0 || errorLevel > 4 {
		errorLevel = -1
	}

	newErr := NiceErrors{
		errorMessage:           errorMessage,
		userFacingErrorMessage: userErrorMessage,
		stackTrace:             string(debug.Stack()) + "~~END OF STACKTRACE~~",
		errorType:              errorType,
		errorLevel:             errorLevel,
	}

	fmt.Println(newErr.Error())
	return &newErr
}

func FromErrorFull(error error, supplementaryErrorMessage string, userErrorMessage string, errorType ErrorType, errorLevel errLevel) *NiceErrors {
	if err, ok := error.(*NiceErrors); ok {
		return err
	}

	if errorLevel < 0 || errorLevel > 4 {
		errorLevel = -1
	}

	errMessage := error.Error()
	if len(strings.TrimSpace(supplementaryErrorMessage)) != 0 {
		errMessage += " --- " + supplementaryErrorMessage
	}

	newErr := NiceErrors{
		errorMessage:           error.Error(),
		userFacingErrorMessage: userErrorMessage,
		stackTrace:             string(debug.Stack()) + "~~END OF STACKTRACE~~",
		errorType:              errorType,
		errorLevel:             errorLevel,
	}

	fmt.Println(newErr.Error())
	return &newErr
}

func FromError(error error) *NiceErrors {
	if err, ok := error.(*NiceErrors); ok {
		return err
	}

	newErr := NiceErrors{
		errorMessage:           error.Error(),
		userFacingErrorMessage: "No error message given",
		errorType:              defaultErrorType,
		errorLevel:             -1,
	}

	fmt.Println(newErr.Error())
	return &newErr
}

func (e *NiceErrors) Error() string {
	return "- User facing error message: " + e.userFacingErrorMessage + "\n" + "- Error message: " + e.errorMessage + "\n" + "- Error level: " + strconv.Itoa(int(e.errorLevel)) + "\n" + "- Stacktrace:\n" + e.stackTrace
}

func (e *NiceErrors) ErrorMessage() string {
	return e.errorMessage
}

func (e *NiceErrors) UserFacingErrorMessage() string {
	return e.userFacingErrorMessage
}

func (e *NiceErrors) StackTrace() string {
	return e.stackTrace
}

func (e *NiceErrors) ErrorType() ErrorType {
	return e.errorType
}

func (e *NiceErrors) ErrorLevel() errLevel {
	return e.errorLevel
}

func (e *NiceErrors) ToJson() string {
	jsonStruct := niceErrorsJson{
		UserFacingErrorMessage: e.userFacingErrorMessage,
	}

	jsonStr, _ := json.Marshal(jsonStruct)
	return string(jsonStr)
}
