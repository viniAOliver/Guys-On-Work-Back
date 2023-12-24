package main

// Imports
import (
	"guys_on_work_back/routes"

	"github.com/gin-gonic/gin"
)

// Initializing project controllers
var (
	userSystemRoutes routes.UserSystemRoutes = routes.NewUserSystemRoutes()
)

// Main function
func main() {

	// Declaring the server variable
	server := gin.Default()

	// Route Group for the User System
	userSystemRoutes.UserSystemRoutes(server)

	// Starting the server, specifying the port on which the service will be available
	server.Run(":8080")

}
