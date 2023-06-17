package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	DB *gorm.DB
}

type Ranking struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Phone   string `gorm:"unique;not null"`
	Current int    `gorm:"not null"`
}

func NewDB() (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to DB")

	app := &Database{
		DB: db,
	}

	return app, nil
}

func (db *Database) ExistsPhone(phone string) bool {
	var count int64
	db.DB.Model(&Ranking{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func (db *Database) Create(userRanking *Ranking) error {
	result := db.DB.Create(userRanking)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *Database) GetByPhone(phone string) (*Ranking, error) {
	var userRanking Ranking
	result := db.DB.Where("phone = ?", phone).First(&userRanking)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userRanking, nil
}

func (db *Database) UpdateCurrent(userRanking Ranking) error {
	result := db.DB.Save(&userRanking)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (db *Database) DeleteAll() error {
	result := db.DB.Delete(&Ranking{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Database) GetAll() (*[]Ranking, error) {
	var ranking []Ranking

	result := db.DB.Order("current").Find(&ranking)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ranking, nil
}

func (db *Database) AutoMigrateTables() error {
	err := db.DB.AutoMigrate(&Ranking{})
	if err != nil {
		return err
	}

	return nil
}
