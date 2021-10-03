package repository

import (
	"gorm.io/gorm"
)

type UOW struct {
	DB        *gorm.DB
	readonly  bool
	committed bool
}

func NewUnitOfWork(db *gorm.DB, flag bool) *UOW {
	if flag {
		return &UOW{
			DB:       db,
			readonly: flag,
		}
	}
	return &UOW{
		DB:       db.Begin(),
		readonly: flag,
	}
}

func (uow *UOW) Commit() {
	if uow.readonly != true {
		uow.DB.Commit()
	}
	uow.committed = true
}

func (uow *UOW) RollBack() {
	if !uow.committed && !uow.readonly {
		uow.DB.Rollback()
	}
}
