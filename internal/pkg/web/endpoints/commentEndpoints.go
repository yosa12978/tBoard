package endpoints

import (
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("postId"))
	err = ce.db.Where("id = ?", ctx.Param("postId")).
		First(&models.Post{}).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	ce.db.Create(&models.Comment{
		Content:  comment.Content,
		PostId:   uint(id),
		AuthorId: ctx.MustGet("account").(models.Account).ID,
	})
	ctx.JSON(201, gin.H{
		"status":  201,
		"message": "created",
	})
}

func (ce *commentEndpoints) Delete(ctx *gin.Context) {
	var obj models.Comment
	err := ce.db.Where("author_id = ? AND id = ?",
		ctx.MustGet("account").(models.Account).ID,
		ctx.Param("id")).First(&obj).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	ce.db.Delete(obj)
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
	})
}
