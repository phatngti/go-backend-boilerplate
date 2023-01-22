package models

import (
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

func (m Models) GetModels() *Models {
	return &m
}
