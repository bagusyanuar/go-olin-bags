package migrations

import (
	"gorm.io/gorm"
)

func prepareTable() []interface{} {
	return []interface{}{
		&User{},
		&Province{},
		&City{},
		&ProductionHouse{},
		&Agent{},
		&SewingAgent{},
		&PrintingAgent{},
		&Material{},
		&Item{},
	}
}
func Migrate(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Province{})
	// db.AutoMigrate(&City{})
	// db.AutoMigrate(&ProductionHouse{})
	// db.AutoMigrate(&Agent{})
	tables := prepareTable()
	db.AutoMigrate(tables...)
}

func Drop(db *gorm.DB) {
	tables := prepareTable()
	db.Migrator().DropTable(tables...)
}
