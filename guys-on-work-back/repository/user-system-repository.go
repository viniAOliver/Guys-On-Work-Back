package repository

// Imports
import (
	"github.com/jinzhu/gorm"
	"guys_on_work_back/entity"
)

// Interface represents to User System Repository
type UserSystemRepository interface {

	// Methods of repository
	CloseConnectionDB()
	UserSystemList() []entity.UserSystem
}

// Struct representing User System Repository
type userSystemRepository struct {

	// Parameters of User System Repository
	db *gorm.DB
}

// Method initial implementation of User System Repository
func NewUserSystemRepository() UserSystemRepository {

	// Command the access the database
	base := "host=localhost user=golang password=root dbname=golang port=5432 sslmode=disable"

	// Connection to database
	db, err := gorm.Open("postgres", base)

	// If the error is non-null, hear error
	if err != nil {

		// Error message
		panic("Falha ao se conectar ao banco de dados: " + err.Error())

	}

	// Migrate of entity from database ( "user_system" )
	db.Table("user_system").AutoMigrate(&entity.UserSystem{})

	// Return the user system repository information
	return &userSystemRepository{
		db: db,
	}

}

// Method to close connection to database
func (db *userSystemRepository) CloseConnectionDB() {

	// Close database connection
	err := db.db.Close()

	// If the error is non-null, hear error
	if err != nil {

		// Error message
		panic("Nao foi possível encerrar comunicação com o banco de dados!")

	}

}

// Method to list user system
func (db *userSystemRepository) UserSystemList() []entity.UserSystem {

	// Defining a variable that will store the system users
	var users_system []entity.UserSystem

	// Executes a database query to retrieve records corresponding to the model associated with &users_system
	db.db.Set("gorm:auto_preload", true).Find(&users_system)

	// Return a list of all users
	return users_system

}
