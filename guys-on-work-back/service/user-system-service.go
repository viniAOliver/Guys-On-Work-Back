package service

// Imports
import (
	"guys_on_work_back/entity"
	"guys_on_work_back/repository"
)

// Interface represents to User System Service
type UserSystemService interface {

	// Methods of service
	UserSystemList() []entity.UserSystem
}

// Struct representing User System Service
type userSystemService struct {

	// Parameters of User System Service
	userSystemRepository repository.UserSystemRepository
}

// Method initial implementation of User System Service
func NewUserSystemService(repo repository.UserSystemRepository) UserSystemService {

	// Return the user system service information
	return &userSystemService{
		userSystemRepository: repo,
	}
}

// Method to list user system
func (service *userSystemService) UserSystemList() []entity.UserSystem {

	// Return the list of system users - calling the service.repository
	return service.userSystemRepository.UserSystemList()

}

