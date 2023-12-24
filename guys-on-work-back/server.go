package main

// Imports
import (
	"github.com/gin-gonic/gin"
	"guys_on_work_back/controller"
	"guys_on_work_back/repository"
	"guys_on_work_back/service"
)

// Initializing project controllers
var (
	userSystemRepository repository.UserSystemRepository = repository.NewUserSystemRepository()
	userSystemService    service.UserSystemService       = service.NewUserSystemService(userSystemRepository)
	userSystemController controller.UserSystemController = controller.NewUserSystemController(userSystemService)
)

// Main function
func main() {

	// Declaring the server variable
	server := gin.Default()

	// Route Group for the User System
	userSystemRoute := server.Group("/user_system")
	{
		userSystemRoute.GET("/", userSystemController.UserSystemList)
		userSystemRoute.GET("/:id", userSystemController.UserSystemDetail)
		userSystemRoute.POST("/create", userSystemController.UserSystemCreate)
		userSystemRoute.PUT("/update/:id", userSystemController.UserSystemUpdate)
		userSystemRoute.DELETE("/delete/:id", userSystemController.UserSystemDelete)
	}

	// Starting the server, specifying the port on which the service will be available
	server.Run(":8080")

}
