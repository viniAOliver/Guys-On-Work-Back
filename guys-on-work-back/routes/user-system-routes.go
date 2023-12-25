package routes

// Imports of packages
import (
	"github.com/gin-gonic/gin"
	"guys_on_work_back/controller"
	"guys_on_work_back/repository"
	"guys_on_work_back/service"
)

// Interface represents to User System Routes
type UserSystemRoutes interface {
	// Methods of routes
	UserSystemRoutes(server *gin.Engine)
}

type userSystemRoutesImpl struct {
	userSystemController controller.UserSystemController
}

// Initializing project controllers
var (
	userSystemRepository repository.UserSystemRepository = repository.NewUserSystemRepository()
	userSystemService    service.UserSystemService       = service.NewUserSystemService(userSystemRepository)
	userSystemController controller.UserSystemController = controller.NewUserSystemController(userSystemService)
)

func NewUserSystemRoutes() UserSystemRoutes {
	return &userSystemRoutesImpl{
		userSystemController: userSystemController,
	}
}

func (r *userSystemRoutesImpl) UserSystemRoutes(server *gin.Engine) {
	userSystemRoute := server.Group("/user_system")
	{
		userSystemRoute.GET("/", r.userSystemController.UserSystemList)
		userSystemRoute.GET("/:id", userSystemController.UserSystemDetail)
		userSystemRoute.POST("/create", userSystemController.UserSystemCreate)
		userSystemRoute.PUT("/update/:id", userSystemController.UserSystemUpdate)
		userSystemRoute.DELETE("/delete/:id", userSystemController.UserSystemDelete)
	}
}
