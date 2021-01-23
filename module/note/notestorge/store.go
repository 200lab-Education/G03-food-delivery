package notestorge

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *store {
	return &store{db: db}
}
