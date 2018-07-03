package pgoptions

import (
	"log"

	"medium-api/config"

	"github.com/go-pg/pg"
)

func GetPgOptions() *pg.Options {
	pgOptions, err := pg.ParseURL(config.GetConfig().DATABASE_URL)

	if err != nil {
		log.Fatal(err)
	}

	return pgOptions
}
