package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/jordancabral/golang-dojo/app/model"
	"github.com/jordancabral/golang-dojo/app/repository"
	"github.com/softbrewery/gojoi/pkg/joi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateConfig ...
func CreateConfig(c *gin.Context) {
	mockTest := Mock{}
	c.ShouldBind(&mockTest)

	validate := validateMock(mockTest)
	if nil != validate {
		fmt.Println("mock mal armado")
		c.JSON(400, gin.H{"message": "mock mal armado"})
	}

	err := repository.CreateConfig(mockTest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating"})
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

// GetAllConfigs ...
func GetAllConfigs(c *gin.Context) {
	mockConfigs := repository.GetAllConfigs()
	c.JSON(200, mockConfigs)
	return
}

// GetConfig ...
func GetConfig(c *gin.Context) {
	id := c.Param("id")
	mockConfig, err := repository.GetConfigById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(200, mockConfig)
	return
}

// RemoveConfig ...
func RemoveConfig(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteConfigByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "path not found"})
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}

// UpdateConfig ...
func UpdateConfig(c *gin.Context) {
	updatedMock := Mock{}
	c.ShouldBind(&updatedMock)
	validate := validateMock(updatedMock)

	if nil != validate {
		fmt.Println("mock mal armado")
		c.JSON(400, gin.H{"message": "mock mal armado"})
	}

	id := c.Param("id")
	idPrimitive, _err := primitive.ObjectIDFromHex(id)
	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating"})
		return
	}

	updatedMock.SetID(idPrimitive)
	err := repository.UpdateConfig(updatedMock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating"})
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}

func validateMock(mock Mock) error {

	schemaCustomHeader := joi.Struct().Keys(joi.StructKeys{
		"key":   joi.String().NonZero(),
		"value": joi.String().NonZero(),
	})

	schemaMock := joi.Struct().Keys(joi.StructKeys{
		"Path":      joi.String().NonZero(),
		"Method":    joi.String().NonZero(),
		"Response":  joi.String().NonZero(),
		"Code":      joi.Int().NonZero(),
		"Headers":   joi.Slice().Items(schemaCustomHeader),
		"ProxyMode": joi.Bool().Required(),
		"ProxyUrl":  joi.String().NonZero(),
	})
	err := joi.Validate(mock, schemaMock)
	return err
}
