package models

import (
	"log"
	core "phatngti/boilerplate/core/crud"
	"phatngti/boilerplate/database"
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	Name      string         `gorm:"column:name" json:"name,omitempty"`
	Email     string         `gorm:"column:email" json:"email,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// Hooks

func (u *UserEntity) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().UTC()
	u.CreatedAt = &now
	u.UpdatedAt = &now
	return nil
}

func (u *UserEntity) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().UTC()
	u.UpdatedAt = &now
	return nil
}

func (u *UserEntity) BeforeDelete(tx *gorm.DB) error {
	now := time.Now().UTC()
	err := tx.Model(&UserEntity{}).Where("id = ?", u.ID).Update("updated_at", &now).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPSQLUserRepository(db *database.Database) (*core.PSqlRepository[UserEntity], error) {
	db.GetPSqlDB().AutoMigrate(&UserEntity{})
	userRepository, err := core.GetCRUDFactory[UserEntity](db, "postgres")
	if err != nil {
		log.Fatal("Failted to create user repository")
		return nil, err
	}
	return userRepository.(*core.PSqlRepository[UserEntity]), nil
}
