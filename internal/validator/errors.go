package validator

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Errors struct {
	errors []gojsonschema.ResultError
}

func (err *Errors) Error() string {
	errors := ""
	for _, desc := range err.errors {
		errors = fmt.Sprintf("%s\n%s", errors, desc)
	}
	return errors
}
