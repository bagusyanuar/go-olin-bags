package migrations

import (
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID       uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email    string         `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password *string        `gorm:"type:text" json:"password"`
	Roles    datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	common.WithTimestampsModel
}

type Province struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Code string    `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	Name string    `gorm:"column:name;type:varchar(255);" json:"name"`
	common.WithTimestampsModel
}

type City struct {
	ID         uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProvinceID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_province_id;not null" json:"province_id"`
	Code       string    `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	Name       string    `gorm:"column:name;type:varchar(255);" json:"name"`
	Province   Province  `gorm:"foreignKey:ProvinceID" json:"province"`
	common.WithTimestampsModel
}

type ProductionHouse struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_user_id;not null;" json:"user_id"`
	CityID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_city_id;not null" json:"city_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone     string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Latitude  float64   `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(11,8)" json:"longitude"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	City      City      `gorm:"foreignKey:CityID" json:"city"`
	common.WithTimestampsModel
}

type Agent struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_user_id;not null;" json:"user_id"`
	CityID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_city_id;not null" json:"city_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone     string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Latitude  float64   `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(11,8)" json:"longitude"`
	Balance   float64   `gorm:"type:double(20,2);default:0;" json:"balance"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	City      City      `gorm:"foreignKey:CityID" json:"city"`
	common.WithTimestampsModel
}

type SewingAgent struct {
	ID                uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionHouseID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_house_id;not null;" json:"production_house_id"`
	Name              string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone             string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address           string    `gorm:"type:text" json:"address"`
	common.WithTimestampsModel
	ProductionHouse ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
}

type PrintingAgent struct {
	ID                uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionHouseID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_house_id;not null;" json:"production_house_id"`
	Name              string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone             string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address           string    `gorm:"type:text" json:"address"`
	common.WithTimestampsModel
	ProductionHouse ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
}

type Material struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name string    `gorm:"type:varchar(255);not null;" json:"name"`
	common.WithTimestampsModel
}

type Item struct {
	ID          uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	MaterialID  uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_material_id;not null;" json:"material_id"`
	Name        string    `gorm:"type:varchar(255);not null;" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       int64     `gorm:"type:bigint(20);default=0" json:"price"`
	common.WithTimestampsModel
	Material Material `gorm:"foreignKey:MaterialID" json:"material"`
}
