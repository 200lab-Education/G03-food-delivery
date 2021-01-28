package notemodel

import "demo/common"

type CreateNote struct {
	common.SQLModel `json:",inline"`
	Title           *string `json:"title" form:"title" gorm:"column:title"`
	Content         *string `json:"content" form:"content" gorm:"column:content"`
}

func (CreateNote) TableName() string {
	return Note{}.TableName()
}
