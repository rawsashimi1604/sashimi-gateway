package main

import (
	"context"
	"fmt"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/logger"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/rproxy"
	"github.com/rs/zerolog/log"
)

func main() {

	logger.SetupLogger()

	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	conn.Ping(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM test")
	if err != nil {
		log.Info().Msg(err.Error())
		log.Fatal().Msg("unable to query rows.")
	}

	for rows.Next() {
		var id int
		var testString string
		if err := rows.Scan(&id, &testString); err != nil {
			log.Fatal().Msg("Error: " + err.Error())
		}
		log.Info().Msg(fmt.Sprintf("%v: %v", id, testString))
	}

	rows.Close()

	api.NewRouter()
	rproxy.ReverseProxy()
	log.Info().Msg("Hello world from admin api.")

}
