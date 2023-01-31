package core_crud_test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	core "phatngti/boilerplate/core/crud"
	"phatngti/boilerplate/database"
	"testing"
	"time"

	"gorm.io/gorm"
)

type ProductEntity struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	Name      string         `gorm:"column:name" json:"name,omitempty"`
	Weight    uint           `gorm:"column:weight" json:"weight,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// Hooks

func (p *ProductEntity) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().UTC()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return nil
}

func (p *ProductEntity) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().UTC()
	p.UpdatedAt = &now
	return nil
}

func (p *ProductEntity) BeforeDelete(tx *gorm.DB) error {
	now := time.Now().UTC()
	err := tx.Model(&ProductEntity{}).Where("id = ?", p.ID).Update("updated_at", &now).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateRepository[E any](db *gorm.DB) *core.PSqlRepository[E] {
	repo := core.NewPSQLRepository[E](db)
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

func getDB() *gorm.DB {
	database := new(database.Database)
	dns := "postgresql://postgres:123456@127.0.0.1:5432/postgres"
	database.InitPSql(dns)
	db := database.GetPSqlDB()
	return db
}
func TestMain(m *testing.M) {
	db := getDB()
	err := db.AutoMigrate(ProductEntity{})
	if err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	os.Exit(ret)
}

func TestGetOne(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)
	criteria := CreateCriteria(&ProductEntity{
		Name: "abc",
	})
	data, err := productRepo.GetOne(ctx, criteria)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data.Name)
}

func TestGetOneById(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)

	data, err := productRepo.GetOneById(ctx, 10)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data)
}

func TestGetAll(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)
	datas, err := productRepo.GetAll(ctx, nil)
	if err != nil {
		panic(err)
	}
	t.Log("datas: ", datas)
}

func TestInsert(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)
	product := ProductEntity{
		Name:   "abcxyz",
		Weight: 100,
	}
	t.Log(product)
	data, err := productRepo.Insert(ctx, &product)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data)
}

func TestFindOneAndUpdate(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)
	newData := ProductEntity{
		Name: "phat ngt123",
	}
	criteria := CreateCriteria(&ProductEntity{
		Name: "abcxyz",
	})

	newProduct, err := productRepo.FindOneAndUpdate(ctx, criteria, &newData)
	if err != nil {
		panic(err)
	}
	t.Log("new product: ", newProduct)
}

func TestFindOneAndUpdateOrInsert(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)
	criteria := CreateCriteria(&ProductEntity{
		Name: "test 2023",
	})

	updateOrInsertData := core.UpdateOrInsert[ProductEntity]{
		NewData: &ProductEntity{
			Name:   "test 2023",
			Weight: 1000,
		},
		ReplaceData: &ProductEntity{
			Name: "test updated 2023",
		},
	}

	data, err := productRepo.FindOneAndUpdateOrInsert(ctx, criteria, &updateOrInsertData)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data)

}

func TestDelete(t *testing.T) {
	db := getDB()
	ctx := context.Background()
	productRepo := CreateRepository[ProductEntity](db)

	record, err := productRepo.GetOneById(ctx, 16)
	if err != nil {
		panic(err)
	}

	err = productRepo.Delete(ctx, &record)
	if err != nil {
		panic(err)
	}
}
