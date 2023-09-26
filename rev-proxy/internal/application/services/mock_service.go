package services

import "github.com/rs/zerolog/log"

type MockService struct{}

func (ms *MockService) Execute() {
	log.Info().Msg("executed the service")
}
