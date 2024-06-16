package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"net/http"
)

func printSqlError(ctx *gin.Context, err error) {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code.Name() {
		case "unique_violation":
			log.Println(errorResponse(err))
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
	}
	log.Println(errorResponse(err))
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
}

func printCookieError(ctx *gin.Context, err error) {
	log.Println(errorResponse(err))
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	ctx.Abort()
}
