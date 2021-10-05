package main

import (
	"github.com/mayamika/mai-lev/handbook/internal/app"
	"github.com/mayamika/mai-lev/handbook/internal/postgres"
)

func main() {
	app.New(app.Config{
		HTTP: app.HttpConfig{
			Addr: `:8080`,
		},
		Postgres: postgres.Config{
			Conn: "postgresql://postgres:pass@localhost:5432/handbook?sslmode=disable",
		},
	}).Run()
}
