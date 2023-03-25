package mapper

import (
	"Belajar-Golang-7/dto"
	"Belajar-Golang-7/entity"
	"errors"
	"strconv"
)

func BookDtoToNewEntity(bookDto dto.BookDTO) entity.BookEntity {
	return entity.BookEntity{
		Title:       bookDto.Title,
		Author:      bookDto.Author,
		Description: bookDto.Desc,
	}
}

func BookDtoToEntity(bookDto dto.BookDTO) (entity.BookEntity, error) {
	bookID, err := strconv.ParseUint(bookDto.Id, 10, 32)
	if err != nil {
		return entity.BookEntity{}, errors.New("ID is not a valid number")
	}

	return entity.BookEntity{
		ID:          uint(bookID),
		Title:       bookDto.Title,
		Author:      bookDto.Author,
		Description: bookDto.Desc,
	}, nil
}

func BookEntityToDto(bookEntity entity.BookEntity) dto.BookDTO {
	return dto.BookDTO{
		Id:     strconv.FormatUint(uint64(bookEntity.ID), 10),
		Title:  bookEntity.Title,
		Author: bookEntity.Author,
		Desc:   bookEntity.Description,
	}
}
