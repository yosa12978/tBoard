package endpoints

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/tBoard/internal/pkg/models"
	"gorm.io/gorm"
)

type PostEndpoints interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	DeleteById(ctx *gin.Context)
}

type postEndpoints struct {
	db *gorm.DB
}

func NewPostEndpoints(db *gorm.DB) PostEndpoints {
	return &postEndpoints{db: db}
}

func (pe *postEndpoints) GetAll(ctx *gin.Context) {
	var posts []models.Post
	pe.db.Find(&posts)
	ctx.JSON(200, posts)
}

func (pe *postEndpoints) GetById(ctx *gin.Context) {
	var post models.Post
	err := pe.db.Model(&models.Post{}).
		Where("id = ?", ctx.Param("id")).
		Preload("Comments").
		First(&post).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"content": "not found",
		})
		return
	}
	ctx.JSON(200, post)
}

func (pe *postEndpoints) Create(ctx *gin.Context) {
	var post models.PostCreateDTO
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(400, gin.H{
			"content": err.Error(),
		})
		return
	}
	pe.db.Create(&models.Post{
		Content:   post.Content,
		Timestamp: uint64(time.Now().Unix()),
		AuthorId:  ctx.MustGet("account").(models.Account).ID,
	})
	ctx.JSON(201, gin.H{
		"content": "created",
	})
}

func (pe *postEndpoints) DeleteById(ctx *gin.Context) {
	if err := pe.db.Delete(&models.Post{}, ctx.Param("id")).Error; err != nil {
		ctx.JSON(400, gin.H{
			"content": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"content": "success",
	})
}
