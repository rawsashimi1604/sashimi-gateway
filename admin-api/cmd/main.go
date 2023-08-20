package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/logger"
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
	router := api.NewRouter()

	log.Info().Msg("starting the admin api.")
	log.Info().Msg("admin api now listening for requests.")

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal().Msg("error when starting the server.")
	}

}
