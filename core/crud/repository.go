package core

import (
	"context"
)

type IRepositoryFactory[T any] interface {
	GetAll(context context.Context, criteria map[string]interface{}) (T, error)
	GetOne()
	GetOneById()
	Insert()
	FindOneOrInsert()
	FindOneAndUpdate()
	FindOneAndUpdateOrInsert()
	FindManyAndUpdate()
	FindManyAndUpdateOrInsert()
	SortDelete()
	HardDelete()
}

// func GetCRUDsFactory[T any](db *database.Database, model T, name string)(IRepositoryFactory[T], error) {
// 	if name == "postgres" {
// 		return &PSqlRepository[RepositoryModel[T], T]{db: db.GetPSqlDB()}, nil
// 	}
// 	return nil, fmt.Errorf("wrong branch name passed")
// }
