package core

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"phatngti/boilerplate/database"
	"testing"
	"time"

	"gorm.io/gorm"
)

type ProductEntity struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty"`
	Name 			string `gorm:"column:name" json:"name,omitempty"`
	Weight 		uint `gorm:"column:weight" json:"weight,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *sql.NullTime `json:"deleted_at,omitempty"`
}

func (p ProductEntity)AuditEntity(isDeleted bool) ProductEntity {
	date := time.Now().UTC()
	deletedDate := sql.NullTime{
		Time: time.Now(),
		Valid: false,
	}
	if isDeleted {
		deletedDate = sql.NullTime{
			Time: date,
			Valid: true,
		}
	}
	p.CreatedAt = &date
	p.UpdatedAt = &date
	p.DeletedAt = &deletedDate
	return p
}

func CreateRepository[E any](db *gorm.DB) *PSqlRepository[E] {
	repo := NewRepository[E](db)
	return repo
}

func MarshalEntity[E any](e *E) map[string]interface{} {
	var inInf map[string]interface{}
	eMarshal, _ := json.Marshal(e)
	json.Unmarshal(eMarshal, &inInf)
	return inInf
}

func CreateCriteria[E any](entity *E) map[string]interface{} {
	return MarshalEntity(entity)
}



func getDB() (*gorm.DB) {
	database := new(database.Database)
	dns := "postgresql://postgres:123456@127.0.0.1:5432/postgres"
	database.InitPSql(dns)
	db := database.GetPSqlDB()
	return db
}
func TestMain(m *testing.M){
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
	data, err := productRepo.GetOne(ctx,criteria)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data.Name)
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
		Name: "abcxyz",
		Weight: 100,
	}.AuditEntity(true)
	t.Log(product)
	data, err := productRepo.Insert(ctx, &product)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data)
}

// func TestFindOneAndUpdate(t *testing.T) {
// 	db := getDB()
// 	ctx := context.Background()
// 	productRepo := NewRepository[ProductEntity, ProductEntity](db)
// 	newData := ProductEntity{
// 		Name: "phat ngt123",
// 	}
// 	criteria := CreateCriteria(&ProductEntity{
// 		Name: "abc",
// 	})

// 	newProduct, err := productRepo.FindOneAndUpdate(ctx, criteria, &newData)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.Log("new product: ", newProduct)
// }

// func TestFindOneAndUpdateOrInsert(t *testing.T) {
// 	db := getDB()
// 	ctx := context.Background()
// 	productRepo := NewRepository[ProductEntity,ProductEntity](db)
// 	criteria := &ProductCriteria{
// 		name: "abc",
// 	}

// 	newData := ProductEntity{
// 		Name: "Phat Ng",
// 	}

// }
