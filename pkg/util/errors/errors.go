package errors

import (
	"fmt"
	"log"
)

// GetFormattedErrorMessage returns a formatted error message for the given
// errorToHandle and errorMessage.
func GetFormattedErrorMessage(errorToHandle error, errorMessage string) string {
	result := ""
	if errorMessage == "" {
		errorMessage = "Error making API call"
	}
	if errorToHandle != nil {
		result = fmt.Sprintf(errorMessage+": %v", errorToHandle.Error())
	}
	return result
}

// HandleError handles the given error
func HandleError(errorToHandle error, errorMessage string) {
	if errorToHandle != nil {
		log.Println(GetFormattedErrorMessage(errorToHandle, errorMessage))
	}
}
