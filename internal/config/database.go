package config

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "pass"
	dbName   = "avito_test"
	dbDriver = "mysql"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)

	db, err := sql.Open(dbDriver, sqlInfo)
	if err != nil {
		log.Error().Msgf("Error in connection: ", err)
	}

	log.Info().Msg("Connected to database")

	return db
}
