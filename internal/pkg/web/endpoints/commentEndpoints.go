package endpoints

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/tBoard/internal/pkg/models"
	"gorm.io/gorm"
)

type CommentEndpoints interface {
	GetPostComments(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type commentEndpoints struct {
	db *gorm.DB
}

func NewCommentEndpoints(db *gorm.DB) CommentEndpoints {
	return &commentEndpoints{db: db}
}

func (ce *commentEndpoints) GetPostComments(ctx *gin.Context) {
	var comments []models.Comment
	ce.db.Where("post_id = ?", ctx.Param("postId")).Find(&comments)
	ctx.JSON(200, comments)
}

func (ce *commentEndpoints) Create(ctx *gin.Context) {
	var comment models.CommentCreateDTO
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{
			"content": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("postId"))
	err = ce.db.Where("id = ?", ctx.Param("postId")).
		First(&models.Post{}).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"content": err.Error(),
		})
		return
	}
	ce.db.Create(&models.Comment{
		Content:  comment.Content,
		PostId:   uint(id),
		AuthorId: ctx.MustGet("account").(models.Account).ID,
	})
	ctx.JSON(201, gin.H{
		"content": "created",
	})
}

func (ce *commentEndpoints) Delete(ctx *gin.Context) {
	if err := ce.db.Delete(&models.Comment{}, ctx.Param("id")).Error; err != nil {
		ctx.JSON(400, gin.H{
			"content": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"content": "success",
	})
}
