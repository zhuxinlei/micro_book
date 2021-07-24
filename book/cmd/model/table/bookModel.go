package table

import (
	"github.com/jinzhu/gorm"
	"github.com/zhuxinlei/micro_book/book/cmd/model/input"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/book"
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
func (u BookModel) GetBook(user input.Book) (output book.BookInfoReply, err error) {
	var data input.Book
	err = u.db.Debug().Where(user).Take(&data).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	return *data.ToOutput(), err
}

func (u BookModel)GetBooks(ids []int)(data []*book.BookInfoReply,err error)  {
	var temp []input.Book
	err = u.db.Debug().Where("id in (?)",ids).Find(&temp).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	for _,k := range temp{
		data = append(data,k.ToOutput())
	}
	return
}
