package core

import (
	"context"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)
type RepositoryModel [E any] interface {
	ToEntity() E
	FromEntity(entity E) interface{}
}

type PSqlRepository[M RepositoryModel[E], E any] struct {
	db *gorm.DB
}

type UpdateOrInsert[E any] struct {
	newData *E
	replaceData *E
}

func NewRepository[M RepositoryModel[E], E any](db *gorm.DB) *PSqlRepository[M, E] {
	return &PSqlRepository[M,E]{
		db: db,
	}
}


func(r *PSqlRepository[M,E]) GetAll(ctx context.Context, criteria map[string]interface{} ) (E, error) {
	var model M
	err := r.db.WithContext(ctx).Where(criteria).Find(&model).Error
	if err != nil {
		return 	*new(E), err
	}
	return model.ToEntity(), nil
}

func(r *PSqlRepository[M,E]) GetOne(ctx context.Context, criteria map[string]interface{}) (E, error) {
	var model M
	err := r.db.WithContext(ctx).Where(criteria).First(&model).Error
	if err != nil {
		return *new(E), err
	}
	fmt.Println("model: ", model)
	return model.ToEntity(), nil
}

func(r *PSqlRepository[M,E]) GetOneById(ctx context.Context, id uint) (E, error) {
	var model M
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return *new(E), err
	}
	return model.ToEntity(), nil
}

func(r *PSqlRepository[M,E]) Insert(ctx context.Context, entity *E) (error) {
	var raw M
	model := raw.FromEntity(*entity).(M)
	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return err
	}
	*entity = model.ToEntity()
	return nil
}

func(r *PSqlRepository[M,E]) FindOneOrInsert(ctx context.Context, criteria map[string]interface{}, entity *E) (E, error) {
	var model M
	err := r.db.WithContext(ctx).Where(criteria).First(&model).Error
	if err != nil {
		return *new(E), err
	}
	if reflect.ValueOf(model).IsZero() {
		var raw M
		newModel := raw.FromEntity(*entity).(M)
		err := r.db.WithContext(ctx).Create(&newModel).Error
		if err != nil {
			return *new(E), err
		}
		*entity = newModel.ToEntity()
		return newModel.ToEntity(), nil
	}
	return model.ToEntity(), nil
}

func(r *PSqlRepository[M,E]) FindOneAndUpdate(ctx context.Context, criteria map[string]interface{}, entity *E) (E, error) {
	var model M
	var err error
	err = r.db.WithContext(ctx).Where(criteria).First(&model).Error
	if err != nil {
		return *new(E), err
	}

	err = r.db.WithContext(ctx).Model(&model).Updates(entity).Error
	if err != nil {
		return *new(E) ,err
	}
	return model.ToEntity(), nil
}

func(r *PSqlRepository[M,E]) FindOneAndUpdateOrInsert(ctx context.Context, criteria map[string]interface{}, data UpdateOrInsert[E]) (E, error){
	var model M
	var err error
	entity, err := r.GetOne(ctx, criteria)
	model = model.FromEntity(entity).(M)

	if err != nil {
		return *new(E), err
	}

	if reflect.ValueOf(model).IsZero() {
		err = r.Insert(ctx, data.newData)
		if err != nil {
			return *new(E), err
		}
	}

	if reflect.ValueOf(data.replaceData).IsValid() {
		err = r.db.WithContext(ctx).Model(&model).Updates(data.replaceData).Error
		if err != nil {
			return *new(E), fmt.Errorf("failed to updated data")
		}
	}
	return model.ToEntity(), nil
}

func(r *PSqlRepository[M,E]) FindManyAndUpdate() {

}

func(r *PSqlRepository[M,E]) FindManyAndUpdateOrInsert() {

}

func(r *PSqlRepository[M,E]) SortDelete() {

}

func(r *PSqlRepository[M,E]) HardDelete() {

}
