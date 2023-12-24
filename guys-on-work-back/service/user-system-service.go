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
	UserSystemDetail(userSystemID string) (*entity.UserSystem, error)
	UserSystemCreate(userSystem entity.UserSystem) (entity.UserSystem, error)
	UserSystemUpdate(userSystemID string, userSystem entity.UserSystem) (entity.UserSystem, error)
	UserSystemDelete(userSystemID string) (entity.UserSystem, error)
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

// Method to detail user system
func (service *userSystemService) UserSystemDetail(userSystemID string) (*entity.UserSystem, error) {

	// Return the data of the system user - calling the service.repository
	return service.userSystemRepository.UserSystemDetail(userSystemID)

}

// Method to create a user system
func (service *userSystemService) UserSystemCreate(userSystem entity.UserSystem) (entity.UserSystem, error) {

	// Return the data of the system user created - calling the service.repository
	return service.userSystemRepository.UserSystemCreate(userSystem)

}

// Method to update a user system
func (service *userSystemService) UserSystemUpdate(userSystemID string, userSystem entity.UserSystem) (entity.UserSystem, error) {

	// Return the data of the system user updated - calling the service.repository
	return service.userSystemRepository.UserSystemUpdate(userSystemID, userSystem)

}

// Method to delete a user system
func (service *userSystemService) UserSystemDelete(userSystemID string) (entity.UserSystem, error) {

	// Return the data of the system user deleted - calling the service.repository
	return service.userSystemRepository.UserSystemDelete(userSystemID)

}