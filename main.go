package main

import (
	categorycontroller "finance-tracker-api/Controllers/CategoryController"
	financecontroller "finance-tracker-api/Controllers/FinanceController"
	usercontroller "finance-tracker-api/Controllers/UserController"
	database "finance-tracker-api/Database"
	helpers "finance-tracker-api/Helpers"
	middleware "finance-tracker-api/Middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
	database.InitDB()
}

func main() {
	router := gin.Default()
	apiGroup := router.Group("/api")

	PrefixAuthGroup := apiGroup.Group("/auth")
	{
		PrefixAuthGroup.POST("/signin", usercontroller.SignIn)
		PrefixAuthGroup.POST("/signup", usercontroller.SignUp)
	}

	AuthGroup := apiGroup.Group("")
	{
		AuthGroup.Use(middleware.Authentication)
		AuthGroup.GET("/auth/user", usercontroller.Profile)

		AuthGroup.GET("/categories", categorycontroller.Index)
		AuthGroup.POST("/categories", categorycontroller.Store)
		AuthGroup.GET("/categories/:id", categorycontroller.Show)
		AuthGroup.PUT("/categories/:id", categorycontroller.Update)
		AuthGroup.DELETE("/categories/:id", categorycontroller.Destroy)

		AuthGroup.GET("/finances", financecontroller.Index)
		AuthGroup.POST("/finances", financecontroller.Store)
		AuthGroup.GET("/finances/:id", financecontroller.Show)
		AuthGroup.PUT("/finances/:id", financecontroller.Update)
		AuthGroup.DELETE("/finances/:id", financecontroller.Destroy)
	}

	router.Run(":" + os.Getenv("PORT"))
}
