package database

import (
	"crud-service-food-recepies/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	var err error
	dsn := "host=localhost user=postgres password=pwdpostgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.Recipe{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Table Created")
}

func GetRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	tx := db.Find(&recipes)
	if tx.Error != nil {
		return []models.Recipe{}, tx.Error
	}

	return recipes, nil
}

func GetRecipe(id string) (models.Recipe, error) {
	var recipe models.Recipe

	tx := db.Where("id = ?", id).First(&recipe)
	if tx.Error != nil {
		return models.Recipe{}, tx.Error
	}

	return recipe, nil
}

func CreateRecipe(recipe models.Recipe) error {
	tx := db.Create(&recipe)

	return tx.Error
}

func UpdateRecipe(recipe models.Recipe) error {
	tx := db.Save(&recipe)

	return tx.Error
}

func DeleteRecipe(id string) error {
	tx := db.Where("id = ?", id).Unscoped().Delete(&models.Recipe{})

	return tx.Error
}
