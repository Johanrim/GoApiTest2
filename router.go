package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})
	r.Use(corsConfig)

	r.GET("/")
	book := r.Group("/book")
	category := r.Group("/company")

	SetUpBookRouter(book)
	SetUpCategoryRouter(category)
}

func SetUpBookRouter(group *gin.RouterGroup) {
	group.GET("")        //list
	group.GET("/:id")    //detail
	group.POST("")       // create
	group.PUT("/:id")    //update
	group.DELETE("/:id") //delete
}

func SetUpCategoryRouter(group *gin.RouterGroup) {

}
