package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/tBoard/internal/pkg/helpers"
	"github.com/yosa12978/tBoard/internal/pkg/models"
	"gorm.io/gorm"
)

type AccountEndpoints interface {
	Login(ctx *gin.Context)
	Signup(ctx *gin.Context)
}

type accountEndpoints struct {
	db *gorm.DB
}

func NewAccountEndpoints(db *gorm.DB) AccountEndpoints {
	return &accountEndpoints{db: db}
}

func (ae *accountEndpoints) Login(ctx *gin.Context) {
	var body models.LoginDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}
	phash, _ := helpers.HashPassword(body.Password)
	var acc models.Account
	r := ae.db.First(&acc, "username = ?", body.Username)

	if r.Error != nil || helpers.CheckPasswordHash(acc.Password, phash) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": r.Error.Error(),
		})
		return
	}
	token, err := helpers.NewJWT(acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (ae *accountEndpoints) Signup(ctx *gin.Context) {
	var body models.SignupDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	var acc models.Account
	r := ae.db.Where("username = ?", body.Username).Limit(1).Find(&acc)
	if r.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": r.Error.Error(),
		})
		return
	}
	if r.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "username is already in use",
		})
		return
	}
	hash, _ := helpers.HashPassword(body.Password)
	acc = models.Account{
		Username: body.Username,
		Password: hash,
		Role:     models.ROLE_USER,
	}
	ae.db.Create(&acc)
	ctx.JSON(201, gin.H{
		"status":  201,
		"message": "created",
	})
}
