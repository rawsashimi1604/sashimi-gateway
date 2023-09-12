package utils

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func PrintAllHeaders(req *http.Request) {
	for name, values := range req.Header {
		for _, value := range values {
			log.Info().Msg(name + ": " + value)
		}
	}
}
