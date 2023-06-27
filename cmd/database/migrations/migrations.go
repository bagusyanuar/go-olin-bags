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
		&Purchasing{},
		&PurchaseItem{},
	}
}
func Migrate(db *gorm.DB) {
	tables := prepareTable()
	db.AutoMigrate(tables...)
}

func Drop(db *gorm.DB) {
	tables := prepareTable()
	db.Migrator().DropTable(tables...)
}
