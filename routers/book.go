package routers

import (
	"github.com/adityasinghtads/api_go/database"
	"github.com/adityasinghtads/api_go/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	}

	router.GET("/books", api.GetBooks)

	return router
}
