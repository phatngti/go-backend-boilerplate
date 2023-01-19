package core

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

func (p ProductEntity) ToEntity() ProductEntity {
	return p
}

func (p ProductEntity) FromEntity(product ProductEntity) interface{} {
	return ProductEntity{
		Name: product.Name,
		Weight: product.Weight,
		ID: product.ID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		DeletedAt: product.DeletedAt,
	}
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

func CreateRepository[E any](db *gorm.DB) *PSqlRepository[E,E] {
	repo := NewRepository[E,E](db)
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
	productRepo := NewRepository[ProductEntity, ProductEntity](db)
	fmt.Println("productrepo: ", &productRepo)
	criteria := CreateCriteria(&ProductEntity{
		Name: "abc",
	})
	data, err := productRepo.GetOne(ctx,criteria)
	if err != nil {
		panic(err)
	}
	t.Log("data: ", data)
}

// func TestInsert(t *testing.T) {
// 	db := getDB()
// 	ctx := context.Background()
// 	productRepo := NewRepository[ProductEntity, ProductEntity](db)
// 	product := ProductEntity{
// 		Name: "abc",
// 		Weight: 8,
// 	}.AuditEntity(true)
// 	t.Log(product)
// 	err := productRepo.Insert(ctx, &product)
// 	if err != nil {
// 		panic(err)
// 	}
// }

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
