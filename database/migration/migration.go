package migration

import (
	"fmt"
	"go-fiber-project-1/database"
	"go-fiber-project-1/pkg/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.UserEntity{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
