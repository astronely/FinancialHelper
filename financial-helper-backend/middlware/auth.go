package middlware

import (
	util "FinancialHelper/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		}

		claims, errMsg := util.ValidateToken(clientToken)
		if errMsg != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"username": claims.Username, "isAuth": true})
		ctx.Abort()
	}
}
