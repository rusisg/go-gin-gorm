package main

import (
	"crud-gin-gorm/lib"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("server started!")

	routes := gin.Default()
	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, "welcome from server")
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	lib.ErrorPanic(err)
}
