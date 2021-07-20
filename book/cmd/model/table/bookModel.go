package table

import (
	"github.com/jinzhu/gorm"
	"micro_book/book/cmd/model/input"
	"micro_book/book/cmd/model/output"
)

type BookModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) BookModel {
	return BookModel{
		db: db,
	}
}

func (u BookModel) Insert(input input.Book) (err error) {
	err = u.db.Debug().Create(&input).Error
	return

}
func (u BookModel) GetBook(user input.Book) (output output.Book, err error) {
	var data input.Book
	err = u.db.Debug().Where(user).Take(&data).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	return data.ToOutput(), err
}
