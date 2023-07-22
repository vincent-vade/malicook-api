package config

import (
	"database/sql"
	"log"
	"os"
	"recipes-api/infra/sqlc"
)

func NewDb() *sqlc.Queries {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	sqlcCtx := sqlc.New(db)
	return sqlcCtx
}
