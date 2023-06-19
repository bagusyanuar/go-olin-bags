package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Agent{})
	fmt.Println("successfully migrating database")
}
