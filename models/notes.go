package models

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/zuams/simple-reminder/db"
)

type Note struct {
	Id     int    `gorm:"primary_key" json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

func GetNotes() (notes []*Note, err error) {
	db, _ := db.New()
	err = db.Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func PostNote(c echo.Context) (*Note, error) {
	db, _ := db.New()
	m := echo.Map{}
	var note Note

	if err := c.Bind(&m); err != nil {
		return nil, err
	}

	// begin a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// do some database operastions in the transaction (use 'tx' from this point, not 'db')
	err := tx.Create(&Note{Title: m["title"].(string), Text: m["text"].(string)}).Error
	if err != nil {
		errRollback := tx.Rollback().Error
		if errRollback != nil {
			return nil, err
		}
		return nil, err
	}

	// commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	err = db.Find(&note).Error
	if err != nil {
		return nil, err
	}

	return &note, err
}

func PutNote(c echo.Context, id int) (*Note, error) {
	db, _ := db.New()
	m := echo.Map{}
	var note Note

	if err := c.Bind(&m); err != nil {
		return nil, err
	}

	// begin a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// do some database operastions in the transaction (use 'tx' from this point, not 'db')
	err := tx.Model(&note).Where("id = ?", id).Update(&Note{Id: id, Title: m["title"].(string), Text: m["text"].(string)}).Error
	if err != nil {
		errRollback := tx.Rollback().Error
		if errRollback != nil {
			return nil, err
		}
		return nil, err
	}

	// commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	err = db.Find(&note).Error
	if err != nil {
		return nil, err
	}

	return &note, err
}

func DeleteNote(c echo.Context, id int) error {
	db, _ := db.New()

	del := db.Where("id = ?", id).Delete(&Note{Id: id})

	if del.RowsAffected < 1 {
		return errors.New("record not found")
	}

	return nil
}
