package connection

import (
	"Kerncheh/medium-api/db/pgoptions"

	"github.com/go-pg/pg"
)

var (
	DBCon *pg.DB
)

func Connect() *pg.DB {
	DBCon = pg.Connect(pgoptions.GetPgOptions())
	return DBCon
}
