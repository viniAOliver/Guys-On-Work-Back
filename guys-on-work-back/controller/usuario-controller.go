package controller

// Imports
import (
	"github.com/gin-gonic/gin"
	"guys_on_work_back/entity"
	"guys_on_work_back/service"
)

// Interface represents to User System Controller
type UserSystemController interface {

	// Methods of controller
	UserSystemList(ctx *gin.Context)
}

// Struct representing User System Controller
type userSystemController struct {

	// Parameters of User System Controller
	service service.UserSystemService
}

// Method initial implementation of User System Controller
func NewUserSystemController(service service.UserSystemService) UserSystemController {

	// Return the user system controller information
	return &userSystemController{
		service: service,
	}

}

// Method to list user system
func (c *userSystemController) UserSystemList(context *gin.Context) {

	// Defining the variable of type User System, in which the return of the UserSystemList service will be stored
	var users_system []entity.UserSystem = c.service.UserSystemList()

	// Return JSON with data from the system user list and status code 200 - OK
	context.JSON(200, gin.H{"data": users_system})
	return

}
