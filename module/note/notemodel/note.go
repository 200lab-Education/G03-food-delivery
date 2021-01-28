package notemodel

import (
	"demo/common"
)

const EntityName = "Note"

var (
	ErrNoteDeactive = common.NewCustomError(nil, "note is deactive", "ErrNoteDeactive")
)

type Note struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" form:"title" gorm:"column:title"`
	Content         string `json:"content" form:"content" gorm:"column:content"`
}

func (Note) TableName() string {
	return "notes"
}
