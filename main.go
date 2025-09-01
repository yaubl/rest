package main

import (
	"bls/api"
	"bls/logger"
	"flag"

	_ "embed"
)

//go:embed db/sql/schema.sql
var ddl string
var port string

func init() {
	flag.StringVar(&port, "port", ":8080", "port to expose the server to.")
	flag.Parse()
}

func main() {
	srv := api.NewServer(port, ddl)
	logger.Log.Fatalf("%v", srv.ListenAndServe())
}
