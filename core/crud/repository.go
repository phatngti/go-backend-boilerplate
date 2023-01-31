package core_crud

import (
	"context"
	"fmt"
	"phatngti/boilerplate/database"
)

type IRepositoryFactory[E any] interface {
	GetAll(ctx context.Context, criteria map[string]interface{}) ([]E, error)
	GetOne(ctx context.Context, criteria map[string]interface{}) (E, error)
	GetOneById(ctx context.Context, id uint) (E, error)
	Insert(ctx context.Context, entity *E) (E, error)
	FindOneOrInsert(ctx context.Context, criteria map[string]interface{}, entity *E) (E, error)
	UpdateById(ctx context.Context, id uint, entity *E) (E, error)
	Update(ctx context.Context, entity *E, newData *E) (E, error)
	FindOneAndUpdate(ctx context.Context, criteria map[string]interface{}, entity *E) (E, error)
	FindOneAndUpdateOrInsert(ctx context.Context, criteria map[string]interface{}, data *UpdateOrInsert[E]) (E, error)
	Delete(ctx context.Context, entity *E) error
}

func GetCRUDFactory[T any](db *database.Database, name string) (IRepositoryFactory[T], error) {
	if name == "postgres" {
		return &PSqlRepository[T]{db: db.GetPSqlDB()}, nil
	}
	return nil, fmt.Errorf("wrong branch name passed")
}
