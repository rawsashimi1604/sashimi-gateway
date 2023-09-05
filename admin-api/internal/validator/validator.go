package validator

// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

import (
	"errors"
	"fmt"

	v_lib "github.com/go-playground/validator/v10"
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
	err := v.lib.Struct(input)

	errorList := make([]error, 0)
	if err != nil {
		for _, err := range err.(v_lib.ValidationErrors) {
			errorList = append(errorList, formatStructError(err))
		}
	}

	if len(errorList) > 0 {
		return errors.New(concatErrors(errorList))
	}
	return nil
}

func (v *Validator) ValidateSimple(input string, validation string) error {
	err := v.lib.Var(input, validation)
	if err != nil {
		return err
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
