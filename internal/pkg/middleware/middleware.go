package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

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
			ctx.Abort()
			return
		}
		token := strings.Replace(header, "Bearer ", "", 1)
		claims, err := helpers.ParseJWT(token)
		if err != nil {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		var account models.Account
		rdb := db.NewRedisClient()

		accjson, err := rdb.Get(fmt.Sprintf("account:%s", claims.(jwt.MapClaims)["id"])).Result()
		if err == nil {
			json.Unmarshal([]byte(accjson), &account)
			ctx.Set("account", account)
			ctx.Next()
			return
		}
		log.Println("cache is empty")
		err = db.GetDB().First(&account, "id = ?", uint(claims.(jwt.MapClaims)["id"].(float64))).Error
		if err != nil {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		jsn, _ := json.Marshal(&account)
		err = rdb.Set(
			fmt.Sprintf("account:%s", claims.(jwt.MapClaims)["id"]),
			string(jsn),
			5*time.Minute).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"status":  500,
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("account", account)
		ctx.Next()
	}
}
