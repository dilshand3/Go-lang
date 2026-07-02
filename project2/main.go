package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")

		c.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"student": "dilshan",
		})
	})

	router.GET("/seconapi/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	router.POST("/api/post", func(ctx *gin.Context) {
		type UserLoginReq struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password"`
		}

		var data UserLoginReq
		fmt.Println(data, "this is intial data")

		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		fmt.Println(data, "this is bindJson data")

		ctx.JSON(http.StatusOK, gin.H{
			"email":    data.Email,
			"password": data.Password,
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
