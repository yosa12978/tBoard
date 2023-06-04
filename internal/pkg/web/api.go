package api

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/tBoard/internal/pkg/db"
	"github.com/yosa12978/tBoard/internal/pkg/middleware"
	"github.com/yosa12978/tBoard/internal/pkg/web/endpoints"
)

func Listen(listener net.Listener) {
	r := gin.Default()
	db.NewRedisClient()
	db := db.GetDB()
	postEndpoints := endpoints.NewPostEndpoints(db)
	commentEndpoints := endpoints.NewCommentEndpoints(db)
	accountEndpoints := endpoints.NewAccountEndpoints(db)

	v1 := r.Group("/v1")

	authRequired := v1.Group("/")
	authRequired.Use(middleware.Authorized())
	{
		authRequired.POST("/posts", postEndpoints.Create)
		authRequired.DELETE("/posts/:id", postEndpoints.DeleteById)
		authRequired.POST("/comments/:postId", commentEndpoints.Create)
		authRequired.DELETE("/comments/:postId", commentEndpoints.Delete)
		authRequired.GET("/whoami", accountEndpoints.Whoami)
	}

	v1.GET("/posts", postEndpoints.GetAll)
	v1.GET("/posts/:id", postEndpoints.GetById)

	v1.GET("/comments/:postId", commentEndpoints.GetPostComments)

	v1.POST("/accounts/login", accountEndpoints.Login)
	v1.POST("/accounts/signup", accountEndpoints.Signup)

	r.RunListener(listener)
}
