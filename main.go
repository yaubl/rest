package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	"api/db"
	"api/middlewares"
	"api/routes"

	_ "embed"

	_ "github.com/glebarez/go-sqlite"

	"github.com/gin-gonic/gin"
)

func setup(ctx *middlewares.AppContext) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.WithAppContext(ctx))

	// setup actual routers idk.
	routes.RegisterRoutes(router)

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
		fmt.Println(err.Error())
	}

	appCtx := &middlewares.AppContext{
		DB:      db.New(database),
		Context: ctx,
	}

	router := setup(appCtx)
	router.Run(port)
}
