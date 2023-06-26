package seeder

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bagusyanuar/go-olin-bags/cmd/database/migrations"
	"github.com/bagusyanuar/go-olin-bags/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminSeeder(db *gorm.DB) {
	if db.Migrator().HasTable(&migrations.User{}) {
		if err := db.Where("JSON_SEARCH(roles, 'all', 'administrator') IS NOT NULL").First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			hash, err := bcrypt.GenerateFromPassword([]byte("administrator"), 13)
			if err != nil {
				panic("failed to create seeder")
			}
			password := string(hash)

			role, _ := json.Marshal([]string{"administrator"})
			admin := model.User{
				Email:    "administrator@gmail.com",
				Username: "administrator",
				Password: &password,
				Roles:    role,
			}
			if err := db.Create(&admin).Error; err != nil {
				panic("failed to create seeder")
			}
			fmt.Println("success create admin seeder")
		}
	} else {
		fmt.Println("table users doesn't exist")
	}
}
