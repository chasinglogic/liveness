package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func postgresCheck(resource string) error {
	log.Info().
		Str("resource", resource).
		Str("checkType", "postgres").
		Msg("testing resource")

	db, err := sql.Open("postgres", resource)
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("%s is unavailable: %s", resource, err.Error())
	}

	return nil
}
