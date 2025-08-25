package main

import (
	"flag"

	"api/routes/v0"
	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {
	router := gin.Default()

	return router
}

var port string

func init() {
	flag.StringVar(&port, "port", ":8080", "port to expose the server to.")
}

func main() {
	router := setup()

	// setup actual routers idk.
	v0.SetupRouter(router)

	// NOTE: NOT ABANDONED, JUST GTG I'LL FINISH THIS LATER OK?
	router.Run(port)
}
