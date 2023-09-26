package services

import "github.com/rs/zerolog/log"

type DomainService struct{}

func (ds *DomainService) Execute() {
	log.Info().Msg("executed the domain service")
}
