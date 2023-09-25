package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadexe/gin-source-code/middleware"
	"github.com/ahmadexe/gin-source-code/models"
	"github.com/gin-gonic/gin"

	"strconv"
)

var users []models.User

func init() {
	fmt.Println("Server is starting...")
	users = []models.User{{ID: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@gmail.com", Age: 20}}
}

func defaultRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func addUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
		
	users = append(users, user)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func getUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	var user models.User
	for _, u := range users {
		if u.ID == id {
			user = u
			break
		}
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func getAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    users,
	})
}

func updateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	var user models.User
	for i, u := range users {
		if u.ID == id {
			user = u
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func main() {
	router := gin.Default()
	router.Use(middleware.Auth())

	router.GET("/", defaultRoute)
	router.POST("/users", addUser)
	router.GET("/users", getAllUsers)
	router.GET("/users/:id", getUser)
	router.PUT("/users", updateUser)
	router.DELETE("/users/:id", delete)
	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}
