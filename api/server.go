package api

import (
	"bls/api/middlewares"
	"bls/db"
	"bls/logger"
	"context"
	"database/sql"
	"net/http"

	"github.com/elisiei/zlog"
	_ "modernc.org/sqlite"
)

type Server struct {
	Port     string
	Database *sql.DB
	Router   *http.ServeMux
}

var queries *db.Queries

func InitDatabase(ddl string) *db.Queries {
	ctx := context.Background()

	database, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	if _, err := database.ExecContext(ctx, ddl); err != nil {
		logger.Log.Warn(err.Error())
	}

	return db.New(database)
}

func NewServer(port string, ddl string) *http.Server {
	queries = InitDatabase(ddl)
	router := NewRouter()
	handler := middlewares.LoggerMiddleware(router)

	logger.Log.Infow("started router.", zlog.F{"port": port, "host": "127.0.0.1"})

	return &http.Server{
		Addr:    port,
		Handler: handler,
	}
}
