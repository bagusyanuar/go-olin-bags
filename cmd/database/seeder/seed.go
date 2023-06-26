package seeder

import "gorm.io/gorm"

func Seed(db *gorm.DB) {
	AdminSeeder(db)
}
