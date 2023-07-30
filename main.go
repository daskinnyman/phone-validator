package main

import (
	"fmt"
	"os"

	"phone_validator/pkg/controllers"
	"phone_validator/pkg/repositories"
	"phone_validator/pkg/serivces"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Step 1: Create db connection instance and return it or error
	dsn := "root:root@tcp(127.0.0.1:3306)/phone_validator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := setUpDbConnection(dsn)

	if err != nil {
		return err
	}

	phoneRepo := repositories.CreatePhoneNumberRepository(db)

	// Inject repo into services
	phoneValidatorService := serivces.CreatePhoneValidatorService(phoneRepo)

	// Init gin router
	router := gin.Default()

	// Set up cors
	router.Use(cors.Default())

	// Inject services into controllers
	controllers.CreatePhoneValidatorHandler(router, phoneValidatorService)

	// run the server
	err = router.Run(":3000")
	if err != nil {
		fmt.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}

func setUpDbConnection(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created connection to database")

	return db, nil
}
