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
	Title           string        `json:"title" form:"title" gorm:"column:title"`
	Content         string        `json:"content" form:"content" gorm:"column:content"`
	Cover           *common.Image `json:"cover" gorm:"column:cover;"`
	Photos          common.Images `json:"photos" gorm:"column:photos;"`
}

func (Note) TableName() string {
	return "notes"
}
