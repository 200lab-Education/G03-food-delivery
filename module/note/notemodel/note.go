package notemodel

import (
	"demo/common"
)

const EntityName = "Note"

type Note struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" form:"title" gorm:"column:title"`
	Content         string `json:"content" form:"content" gorm:"column:content"`
}

func (Note) TableName() string {
	return "notes"
}
