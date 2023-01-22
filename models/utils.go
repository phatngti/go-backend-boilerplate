package models

import (
	"encoding/json"
	core "phatngti/boilerplate/core/crud"

	"gorm.io/gorm"
)

func CreatePSQLRepository[E any](db *gorm.DB) *core.PSqlRepository[E] {
	repo := core.NewRepository[E](db)
	return repo
}

func MarshalJSON[E any](e *E) map[string]interface{} {
	var inInf map[string]interface{}
	eMarshal, _ := json.Marshal(e)
	json.Unmarshal(eMarshal, &inInf)
	return inInf
}

func CreateCriteria[E any](entity *E) map[string]interface{} {
	return MarshalJSON(entity)
}
