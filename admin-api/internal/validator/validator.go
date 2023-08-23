package validator

// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

import (
	"errors"
	"fmt"

	v_lib "github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Validator struct {
	lib *v_lib.Validate
}

func NewValidator() *Validator {
	return &Validator{
		lib: v_lib.New(),
	}
}

func (v *Validator) ValidateStruct(input interface{}) error {
	log.Info().Msg("hello world from validator struct")
	err := v.lib.Struct(input)

	errorList := make([]error, 0)
	for _, err := range err.(v_lib.ValidationErrors) {
		errorList = append(errorList, formatStructError(err))
	}

	if len(errorList) > 0 {
		return errors.New(concatErrors(errorList))
	}
	return nil
}

func concatErrors(errors []error) string {
	errorStr := ""
	for _, err := range errors {
		errorStr += err.Error()
	}
	return errorStr
}

func formatStructError(structError v_lib.FieldError) error {
	// namespace is tag, value: <value>
	return fmt.Errorf("%v is %v, value: %v", structError.Namespace(), structError.Tag(), structError.Value())
}
