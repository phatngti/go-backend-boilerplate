package core

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	RFC3339     = "2006-01-02T15:04:05Z07:00"
)

func MapToBson(data map[string]interface{}) bson.D {
	var compoundIndex bson.D
	for key, value := range data {
		compoundIndex = append(compoundIndex, bson.E{Key: key, Value: value})
	}

	return compoundIndex
}

func StructToJSON[E any](e *E) map[string]interface{} {
	var inInf map[string]interface{}
	eMarshal, _ := json.Marshal(e)
	fmt.Println("eMarshal: ", inInf)
	json.Unmarshal(eMarshal, &inInf)

	return inInf
}

func StructToBSON[E any](e *E) bson.D {
	json := StructToJSON(e)
	return MapToBson(json)
}
