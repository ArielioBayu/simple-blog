package models

import (
	"github.com/gofrs/uuid"
)

type Film struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Author   string `gorm:"type:varchar(100)" json:"author"`
	Post     string `gorm:"type:text" json:"post"`
	Category string `gorm:"type:varchar(100)" json:"category"`
}

type User struct {
	ID       string `gorm:"primaryKey;type:char(36)" json:"id"`
	Username string `gorm:"type:varchar(100)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	RoleID   string `gorm:"type:char(36)" json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"`
}

type Role struct {
	ID   string `gorm:"primaryKey;type:char(36)" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
}

func GenerateUUID() string {
	u, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return u.String()
}
