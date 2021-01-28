package notemodel

import "demo/common"

type CreateNote struct {
	common.SQLModel `json:",inline"`
	Title           *string        `json:"title" form:"title" gorm:"column:title"`
	Content         *string        `json:"content" form:"content" gorm:"column:content"`
	Cover           *common.Image  `json:"cover" gorm:"column:cover;"`
	Photos          *common.Images `json:"photos" gorm:"column:photos;"`
}

func (CreateNote) TableName() string {
	return Note{}.TableName()
}
