package repository

import (
	"fmt"
	"os"

	"github.com/Angieuski/Desafio-Edu/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UiRepository interface {
	Save(ui entity.Ui)
	Update(ui entity.Ui)
	Delete(ui entity.Ui)
	FindAll() []entity.Ui
	CloseDB()
	FindByEmail(email string, user *entity.Ui) error
}

type Database struct {
	connection *gorm.DB
}

func NewRepository() UiRepository {
	db, err := gorm.Open("sqlite3", "uis.db")
	if err != nil {
		fmt.Println("Failed to connect database:", err)
		os.Exit(1)
	}
	db.AutoMigrate(&entity.Ui{})
	return &Database{connection: db}
}

func (db *Database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *Database) Save(ui entity.Ui) {
	db.connection.Create(&ui)
}

func (db *Database) FindByEmail(email string, user *entity.Ui) error {
	return db.connection.Where("email = ?", email).First(user).Error
}

func (db *Database) Update(ui entity.Ui) {
	db.connection.Save(&ui)
}

func (db *Database) Delete(ui entity.Ui) {
	db.connection.Delete(&ui)
}

func (db *Database) FindAll() []entity.Ui {
	var uis []entity.Ui
	db.connection.Set("gorm:auto_preload", true).Find(&uis)
	return uis
}
