package repository

import (
	"errors"
	"fmt"

	"github.com/Kamva/mgm"
	. "github.com/jordancabral/golang-dojo/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "mock_server", options.Client().ApplyURI("mongodb://localhost:27017"))
	if nil != err {
		panic(err)
	}
}

// GetConfig
func GetConfig(method, path string) (*Mock, error) {
	mock := Mock{}
	coll := mgm.Coll(&Mock{})
	err := coll.First(bson.M{"path": path, "method": method}, &mock)
	if err != nil {
		fmt.Println("Path not found")
		return nil, errors.New("path not found")
	}
	return &mock, nil
}

func GetConfigById(id string) (*Mock, error) {
	mock := Mock{}
	coll := mgm.Coll(&Mock{})
	err := coll.FindByID(id, &mock)
	if err != nil {
		fmt.Println("id not found")
		return nil, errors.New("id not found")
	}
	return &mock, nil
}

func GetAllConfigs() Mocks {
	ctx := mgm.Ctx()
	coll := mgm.Coll(&Mock{})
	result, _err := coll.Find(ctx, bson.D{})

	if _err != nil {
		panic(_err)
	}

	mockList := []Mock{}
	result.All(ctx, &mockList)
	fmt.Println("Configs loaded from DB:")
	fmt.Println(mockList)

	return mockList
}

// CreateConfig ...
func CreateConfig(mock Mock) error {
	coll := mgm.Coll(&Mock{})

	err := coll.Create(&mock)
	if err != nil {
		fmt.Println("ERROR:", err)
		return errors.New("Error creating")
	}
	return nil
}

// DeleteConfigByID ...
func DeleteConfigByID(id string) error {
	ctx := mgm.Ctx()
	coll := mgm.Coll(&Mock{})
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex ERROR:", err)
		return errors.New("Error deleting")
	}
	_, err2 := coll.DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err2 != nil {
		fmt.Println("ERROR:", err2)
		return errors.New("Error deleting")
	}
	return nil
}

// UpdateConfig ...
func UpdateConfig(mock Mock) error {
	coll := mgm.Coll(&Mock{})
	err := coll.Update(&mock)
	if err != nil {
		fmt.Println("ERROR:", err)
		return errors.New("Error updating")
	}
	return nil
}
