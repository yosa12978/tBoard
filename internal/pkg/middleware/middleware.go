package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yosa12978/tBoard/internal/pkg/db"
	"github.com/yosa12978/tBoard/internal/pkg/helpers"
	"github.com/yosa12978/tBoard/internal/pkg/models"
)

func Authorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": "unauthorized",
			})
			return
		}
		token := strings.Replace(header, "Bearer ", "", 1)
		claims, err := helpers.ParseJWT(token)
		if err != nil {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": err.Error(),
			})
			return
		}
		var account models.Account
		err = db.GetDB().First(&account, "id = ?", uint(claims.(jwt.MapClaims)["id"].(float64))).Error
		if err != nil {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": err.Error(),
			})
			return
		}
		ctx.Set("account", account)
		ctx.Next()
	}
}
