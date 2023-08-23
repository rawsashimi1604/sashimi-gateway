package validator

// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

import (
	v "github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Validator struct {
	validator *v.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: v.New(),
	}
}

func (v *Validator) Validate() {
	log.Info().Msg("hello world from validator")

}
