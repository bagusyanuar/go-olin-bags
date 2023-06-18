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

type Agent struct {
	ID      uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID  uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_user_id;not null;" json:"user_id"`
	Name    string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone   string    `gorm:"type:varchar(25);" json:"phone"`
	Address string    `gorm:"type:text" json:"address"`
	common.WithTimestampsModel
	User User `gorm:"foreignKey:UserID" json:"user"`
}
