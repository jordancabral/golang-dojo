package repository

import (
	"fmt"

	"github.com/Kamva/mgm"
	. "github.com/jordancabral/golang-dojo/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "mock_server", options.Client().ApplyURI("mongodb://localhost:27017"))
	if nil != err {
		panic(err)
	}
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

func CreateConfig(mock Mock) {
	coll := mgm.Coll(&Mock{})

	_err := coll.Create(&mock)
	if _err != nil {
		panic(_err)
	}
}

func DeleteConfig(mock Mock) {
	coll := mgm.Coll(&Mock{})

	_err := coll.Delete(&mock)
	if _err != nil {
		panic(_err)
	}
}
