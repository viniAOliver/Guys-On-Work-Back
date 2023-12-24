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
	}

	// Starting the server, specifying the port on which the service will be available
	server.Run(":8888")

}
