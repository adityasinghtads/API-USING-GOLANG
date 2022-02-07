package routers

import (
	"github.com/adityasinghtads/api_go/database"
	"github.com/adityasinghtads/api_go/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default() //get the deault engine for further cunstomization 
	api := handler.Handler{
		DB: database.GetDB(), // set the handler DB
	}

	router.GET("/books", api.GetBooks) //set the functiom for this url  http://localhost:8080/books : Get request 
	// while calling handler function, gin will pass *gin.Context in the handler function 
	router.POST("/book",api.SaveBook)
	return router
}
