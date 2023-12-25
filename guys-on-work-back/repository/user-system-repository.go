package repository

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"guys_on_work_back/entity"
)

// Interface represents to User System Repository
type UserSystemRepository interface {

	// Methods of repository
	UserSystemList() []entity.UserSystem
	UserSystemDetail(userSystemID string) (*entity.UserSystem, error)
	UserSystemByEmailDetail(userSystemID string) (*entity.UserSystem, error)
	UserSystemCreate(userSystem entity.UserSystem) (entity.UserSystem, error)
	UserSystemUpdate(userSystemID string, userSystem entity.UserSystem) (entity.UserSystem, error)
	UserSystemDelete(userSystemID string) (entity.UserSystem, error)
}

// Struct representing User System Repository
type userSystemRepository struct {

	// Parameters of User System Repository
	db *gorm.DB
}

// Method initial implementation of User System Repository
func NewUserSystemRepository() UserSystemRepository {

	// Command the access the database
	base := "host=localhost user=root password=root dbname=guys_work_db port=5432 sslmode=disable"

	// Connection to database
	db, err := gorm.Open("postgres", base)

	// If the error is non-null, hear error
	if err != nil {

		// Error message
		panic("Falha ao se conectar ao banco de dados: " + err.Error())

	}

	var nameTable string = "user_system"

	// TODO: Mudar isso para package de migrates
	if db.Table(nameTable).HasTable(&entity.UserSystem{}) {

		// Recuperar dados da tabela
		var users []entity.UserSystem
		db.Table(nameTable).Find(&users)

		// Converter dados para formato desejado (por exemplo, JSON)
		data, err := json.MarshalIndent(users, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		// Salvar dados em um arquivo (por exemplo, users_dump.json)
		fileName := fmt.Sprintf("migrations_dump/%s_dump.json", nameTable)
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Escrever dados no arquivo
		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		db.Table(nameTable).DropTableIfExists(&entity.UserSystem{})

		// Ler dados do arquivo de dump
		file, err = os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var usuarios []entity.UserSystem
		err = json.NewDecoder(file).Decode(&usuarios)
		if err != nil {
			log.Fatal(err)
		}

		// Migrate of entity from database ( "user_system" )
		db.Table(nameTable).AutoMigrate(&entity.UserSystem{})
		// Inserir dados na tabela
		for _, user := range users {
			db.Table(nameTable).Create(&user)
		}
	} else {

		// Migrate of entity from database ( "user_system" )
		db.Table(nameTable).AutoMigrate(&entity.UserSystem{})
	}

	// Return the user system repository information
	return &userSystemRepository{
		db: db,
	}

}

// Method to list user system
func (db *userSystemRepository) UserSystemList() []entity.UserSystem {

	// Defining a variable that will store the system users
	var userSystem []entity.UserSystem

	// Executes a database query to retrieve records corresponding to the model associated with &users_system
	db.db.Table("user_system").Set("gorm:auto_preload", true).Find(&userSystem)

	// Return a list of all users
	return userSystem

}

// Method to detail a user system
func (db *userSystemRepository) UserSystemDetail(userSystemID string) (*entity.UserSystem, error) {

	// Defining a variable that will store the system users
	var userSystem entity.UserSystem

	// Execute the query, and capturing the result and the possible error
	// SELECT * FROM userSystem WHERE ID = userSystemID;
	result := db.db.Table("user_system").Where("id = ?", userSystemID).First(&userSystem)

	// If there is an result.Error
	if result.Error != nil {

		// Return {data: null, error: result.Error}
		return nil, result.Error
	}

	// Return {data: userSystem, error: null}
	return &userSystem, nil

}

// Method to detail a user system by Detail
func (db *userSystemRepository) UserSystemByEmailDetail(userSystemEmail string) (*entity.UserSystem, error) {

	// Defining a variable that will store the system users
	var userSystem entity.UserSystem

	// Execute the query, and capturing the result and the possible error
	// SELECT * FROM userSystem WHERE ID = userSystemID;
	result := db.db.Table("user_system").Where("user_system_email = ?", userSystemEmail).First(&userSystem)

	// If there is an result.Error
	if result.Error != nil {

		// Return {data: null, error: result.Error}
		return nil, result.Error
	}

	// Return {data: userSystem, error: null}
	return &userSystem, nil

}

// Method to create a user system
func (db *userSystemRepository) UserSystemCreate(userSystem entity.UserSystem) (entity.UserSystem, error) {

	// Calling the create system user in the database, and capturing and checking the possible error
	if err := db.db.Table("user_system").Create(&userSystem).Error; err != nil {

		// Return the user system and the error
		return userSystem, err
	}

	// Return the user system and null
	return userSystem, nil

}

// Method to update a user system
func (db *userSystemRepository) UserSystemUpdate(userSystemID string, userSystem entity.UserSystem) (entity.UserSystem, error) {

	// Calling the update a system user in the database, and capturing and checking the possible error
	if err := db.db.Table("user_system").Where("id = ?", userSystemID).Updates(&userSystem).Error; err != nil {

		// Return the user system and the error
		return userSystem, err
	}

	// Return the user system and null
	return userSystem, nil

}

// Method to delete a user system
func (db *userSystemRepository) UserSystemDelete(userSystemID string) (entity.UserSystem, error) {

	// Defining a variable that will store the system users
	var userSystem entity.UserSystem

	// Calling the delete system user in the database, and capturing and checking the possible error
	if err := db.db.Table("user_system").Where("id = ?", userSystemID).Delete(&userSystem).Error; err != nil {

		// Return the user system and the error
		return userSystem, err
	}

	// Return the user system and null
	return userSystem, nil

}
