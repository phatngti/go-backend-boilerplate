package core

import (
	"context"
	"reflect"

	"gorm.io/gorm"
)


type PSqlRepository[E any] struct {
	db *gorm.DB
}

type UpdateOrInsert[E any] struct {
	newData *E
	replaceData *E
}

func NewRepository[ E any](db *gorm.DB) *PSqlRepository[ E] {
	return &PSqlRepository[E]{
		db: db,
	}
}

func(r *PSqlRepository[E]) GetAll(ctx context.Context, criteria map[string]interface{} ) ([]E, error) {
	var model []E
	err := r.db.WithContext(ctx).Where(criteria).Find(&model).Error
	if err != nil {
		return 	*new([]E), err
	}
	return model, nil
}

func(r *PSqlRepository[E]) GetOne(ctx context.Context, criteria map[string]interface{}) (E, error) {
	var model E
	err := r.db.WithContext(ctx).Where(criteria).First(&model).Error
	if err != nil {
		return *new(E), err
	}
	return model, nil
}

func(r *PSqlRepository[E]) GetOneById(ctx context.Context, id uint) (E, error) {
	var model E
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return *new(E), err
	}
	return model, nil
}

func(r *PSqlRepository[E]) Insert(ctx context.Context, entity *E) (E, error) {
	err := r.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return *new(E), err
	}
	return *entity, nil
}

func(r *PSqlRepository[E]) FindOneOrInsert(ctx context.Context, criteria map[string]interface{}, entity *E) (E, error) {
	var model E
	err := r.db.WithContext(ctx).Where(criteria).First(&model).Error
	if err != nil {
		return *new(E), err
	}
	if reflect.ValueOf(model).IsZero() {
		err := r.db.WithContext(ctx).Create(&entity).Error
		if err != nil {
			return *new(E), err
		}
		return *entity, nil
	}
	return model, nil
}

func(r *PSqlRepository[E]) FindOneAndUpdate(ctx context.Context, criteria map[string]interface{}, entity *E) (E, error) {
	var model E
	var err error
	err = r.db.WithContext(ctx).Where(criteria).First(&model).Error
	if err != nil {
		return *new(E), err
	}

	err = r.db.WithContext(ctx).Model(&model).Updates(entity).Error
	if err != nil {
		return *new(E) ,err
	}
	return model, nil
}

func(r *PSqlRepository[E]) FindOneAndUpdateOrInsert(ctx context.Context, criteria map[string]interface{}, data UpdateOrInsert[E]){

}

func(r *PSqlRepository[E]) FindManyAndUpdate() {

}

func(r *PSqlRepository[E]) FindManyAndUpdateOrInsert() {

}

func(r *PSqlRepository[E]) SortDelete() {

}

func(r *PSqlRepository[E]) HardDelete() {

}
