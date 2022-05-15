package sql

import (
	"gorm.io/gorm"
)

type Persister struct {
	db *gorm.DB
}

func NewPersister(db *gorm.DB) *Persister {
	return &Persister{db: db}
}

func (p *Persister) UseKey() (string, error) {
	var (
		key string
		tx  = p.db.Debug().Begin()
		err error
	)
	defer func() {
		r := recover()
		if err != nil || r != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}()

	err = tx.Table("keys").Raw("SELECT `key` FROM `keys` WHERE is_used=0 LIMIT 1 FOR UPDATE").Scan(&key).Error
	if err != nil {
		return "", err
	}

	err = tx.Table("keys").Where("`key` = ?", key).Update("is_used", true).Error
	if err != nil {
		return "", err
	}

	return key, nil
}
