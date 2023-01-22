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

func (r *PSqlRepository[E]) UpdateById(ctx context.Context, id uint, entity *E ) (E, error) {
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(entity).Error
	if err != nil {
		return *new(E), err
	}
	return *entity, nil
}

func (r *PSqlRepository[E]) Update(ctx context.Context, entity *E, newData *E) (E, error) {
	err := r.db.WithContext(ctx).Model(&entity).Updates(newData).Error
	if err != nil {
		return *new(E), err
	}
	return *entity, nil
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

func(r *PSqlRepository[E]) FindOneAndUpdateOrInsert(ctx context.Context, criteria map[string]interface{}, data UpdateOrInsert[E]) (E, error){
	// Get record by criteria
	var record E
	var err error
	record, err = r.GetOne(ctx, criteria)

	if err != nil && err.Error() != RecordNotFound {
		return *new(E), err
	}
	if err != nil && err.Error() == RecordNotFound {
		if data.newData != nil {
			record, err := r.Insert(ctx, data.newData)
			if err != nil {
				return *new(E), err
			}
			return record, nil
		}
	}
	if data.replaceData != nil {
		updateRecord, err := r.Update(ctx, &record, data.replaceData)
		if err != nil {
			return *new(E), err
		}
		return updateRecord, nil
	}
	return record, nil
}

// func(r *PSqlRepository[E]) FindManyAndUpdate() {

// } // Update with batch -> Later

// func(r *PSqlRepository[E]) FindManyAndUpdateOrInsert() {

// } // Update with batch -> Later

func(r *PSqlRepository[E]) Delete(ctx context.Context, entity *E) (error) {
	err := r.db.WithContext(ctx).Delete(entity).Error
	if err != nil {
		return err
	}
	return nil
}
