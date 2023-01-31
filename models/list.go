package models

import (
	"fmt"
	"log"
	core "phatngti/boilerplate/core/crud"
	"phatngti/boilerplate/database"
)

type Models struct {
	user *core.PSqlRepository[UserEntity]
}

func (m *Models) InitModels(db *database.Database) {
	user, err := GetPSQLUserRepository(db)
	if err != nil {
		panic(err)
	}
	m.user = user
}

func (m *Models) GetModels() (*Models, error) {
	if m == nil {
		log.Fatal("Cannot init models")
		return nil, fmt.Errorf("models cannot initilized")
	}
	return m, nil
}
