package main

import (
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"fmt"

	"insurance/router"
	util "insurance/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	//Adding env
	var config util.Config
	initEnv(&config)

	fmt.Printf("APP is running in: %s Host and %s Port  \n", config.Host, config.Port)

	//Connecting to MariaDB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MariaDB connected...")

	//Adding Tables
	err = db.AutoMigrate(&util.User{})
	if err != nil {
		fmt.Println(err)
	}

	// insertProduct := &util.Product{Code: "D42", Price: 100}

	insertUser := &util.User{
		ID:        uuid.New().String(),
		FirstName: "Nehal",
		LastName:  "Singh",
		Email:     "nehal@gmail.com",
		Password:  "djhsdcdvj",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(insertUser)

	fmt.Println("Reached before Token Map")

	router.Router(gin.New())
}

func initEnv(config *util.Config) {
	config.Host = util.ViperEnvVariable("HOST")
	config.Port = util.ViperEnvVariable("PORT")
	config.User = util.ViperEnvVariable("USER")
	config.Pass = util.ViperEnvVariable("PASS")
	config.DB = util.ViperEnvVariable("DB_NAME")
}
