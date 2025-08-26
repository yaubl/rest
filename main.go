package main

import (
	"context"
	"database/sql"
	"flag"
	"log"

	"api/db"
	"api/middlewares"
	"api/routes/v0"
	_ "embed"
	_ "github.com/glebarez/go-sqlite"

	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {
	router := gin.Default()

	return router
}

var port string

//go:embed db/sql/schema.sql
var ddl string

func init() {
	flag.StringVar(&port, "port", ":8080", "port to expose the server to.")
}

func main() {
	ctx := context.Background()

	database, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := database.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	appCtx := &middlewares.AppContext{
		DB:      db.New(database),
		Context: ctx,
	}
	router := setup()
	router.Use(middlewares.WithAppContext(appCtx))

	// setup actual routers idk.
	v0.SetupRouter(router)

	router.Run(port)
}
