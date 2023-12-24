package controller

// Imports
import (
	"github.com/gin-gonic/gin"
	"guys_on_work_back/entity"
	"guys_on_work_back/service"
	"guys_on_work_back/util"
)

// Interface represents to User System Controller
type UserSystemController interface {

	// Methods of controller
	UserSystemList(context *gin.Context)
	UserSystemDetail(context *gin.Context)
	UserSystemCreate(context *gin.Context)
	UserSystemUpdate(context *gin.Context)
	UserSystemDelete(context *gin.Context)
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

// Method to detail a user system
func (c *userSystemController) UserSystemDetail(context *gin.Context) {

	// Extract user ID from the request parameters or JSON body, assuming it's part of the request
	userSystemID := context.Param("id")

	// Checking the user ID is not null
	if userSystemID == "" {

		context.JSON(400, gin.H{"error": "ID é obrigatório!"})
		return

	}

	// Defining the variable of type User System, in which the return of the UserSystemList service will be stored
	userSystem, err := c.service.UserSystemDetail(userSystemID)

	if err != nil {
		context.JSON(400, gin.H{"error": err})
	}

	// Return JSON with data from the system user list and status code 200 - OK
	context.JSON(200, gin.H{"data": userSystem})
	return

}

// Method to create a user system
func (c *userSystemController) UserSystemCreate(context *gin.Context) {

	// Defining a variable that will be filled with the request data
	var userSystem entity.UserSystem

	// Checking whether the information sent is valid according to the entity
	err := context.ShouldBindJSON(&userSystem)

	// If there is an error
	if err != nil {

		// Return JSON with the error found in the system user registration and status code 400 - BAD REQUEST
		context.JSON(400, gin.H{"error": err.Error()})
		return

	}

	// Calling the hash password method using the util, and capturing the hashed password and the possible error
	hashedPassword, err := util.HashPassword(userSystem.UserSystemPassword)

	// If there is an error
	if err != nil {

		// Return JSON with the error found and status code 500 - INTERNAL ERROR
		context.JSON(500, gin.H{"error": err})
		return

	}

	// Atribute the hashed password in the user system
	userSystem.UserSystemPassword = hashedPassword

	// Calling the system user's create method using the service, and capturing the created user and the possible error
	createdUser, err := c.service.UserSystemCreate(userSystem)

	// If there is an error
	if err != nil {

		// Return JSON with the error found in the system user registration and status code 400 - BAD REQUEST
		context.JSON(400, gin.H{"error": err.Error()})
		return

	}

	// Return JSON with the created system user record and status code 201 – CREATED
	context.JSON(201, gin.H{"data": createdUser})
	return

}

// Method to update a user system
func (c *userSystemController) UserSystemUpdate(context *gin.Context) {

	// Extract user ID from the request parameters or JSON body, assuming it's part of the request
	userSystemID := context.Param("id")

	// Checking the user ID is not null
	if userSystemID == "" {

		context.JSON(400, gin.H{"error": "ID é obrigatório!"})
		return

	}

	// Fetch the existing user system based on the user ID
	_, err := c.service.UserSystemDetail(userSystemID)

	// Check if the user exists
	if err != nil {
		context.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Defining a variable that will be filled with the request data
	var updatedUserSystem entity.UserSystem

	// Checking whether the information sent is valid according to the entity
	err = context.ShouldBindJSON(&updatedUserSystem)

	// If there is an error
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Calling the system user's create method using the service, and capturing the created user and the possible error
	updatedUser, err := c.service.UserSystemUpdate(userSystemID, updatedUserSystem)

	// If there is an error
	if err != nil {

		// Return JSON with the error found in the system user registration and status code 400 - BAD REQUEST
		context.JSON(400, gin.H{"error": err.Error()})
		return

	}

	// Return JSON with the updated system user record and status code 200 – OK
	context.JSON(200, gin.H{"data": updatedUser})
	return

}

// Method to delete a user system
func (c *userSystemController) UserSystemDelete(context *gin.Context) {

	// Extract user ID from the request parameters or JSON body, assuming it's part of the request
	userSystemID := context.Param("id")

	// Checking the user ID is not null
	if userSystemID == "" {

		context.JSON(400, gin.H{"error": "ID é obrigatório!"})
		return

	}

	// Calling the system user's create method using the service, and capturing the created user and the possible error
	deleteUser, err := c.service.UserSystemDelete(userSystemID)

	// If there is an error
	if err != nil {

		// Return JSON with the error found in the system user registration and status code 400 - BAD REQUEST
		context.JSON(400, gin.H{"error": err.Error()})
		return

	}

	// Return JSON with the deleted system user record and status code 204 – NO CONTENT
	context.JSON(204, gin.H{"data": deleteUser})
	return

}
