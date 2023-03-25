package service

import (
	"Belajar-Golang-7/db"
	"Belajar-Golang-7/entity"
	"errors"
)

func InsertBook(bookEntity entity.BookEntity) error {
	db := db.GetDB()

	if err := db.Create(&bookEntity).Error; err != nil {
		return err
	}

	return nil
}

func GetAllBook() []entity.BookEntity {
	db := db.GetDB()
	var books []entity.BookEntity

	db.Find(&books)
	return books
}

func GetBookByID(id uint) (entity.BookEntity, error) {
	db := db.GetDB()
	var book entity.BookEntity

	if err := db.First(&book, id).Error; err != nil {
		return book, errors.New("book not found")
	}

	return book, nil
}

func UpdateBook(bookEntity entity.BookEntity) error {
	db := db.GetDB()
	model := entity.BookEntity{}

	result := db.Model(&model).Where("id = ?", bookEntity.ID).Updates(&bookEntity)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

func DeleteBookByID(id uint) error {
	db := db.GetDB()
	model := entity.BookEntity{}
	result := db.Delete(&model, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("book is not exist")
	}

	return nil
}
