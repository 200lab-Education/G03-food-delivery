package notemodel

import "demo/common"

type CreateNote struct {
	common.SQLModel `json:",inline"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	Title           *string        `json:"title" form:"title" gorm:"column:title;"`
	Content         *string        `json:"content" form:"content" gorm:"column:content;"`
	Cover           *common.Image  `json:"json" gorm:"column:cover;"`
	Photos          *common.Images `json:"photos" gorm:"column:photos;"`
	CoverImgId      int            `json:"cover_img_id" gorm:"-"`
}

func (CreateNote) TableName() string {
	return Note{}.TableName()
}

func (n *CreateNote) GetImageIds() []int {
	return []int{n.CoverImgId}
}

func (n *CreateNote) GetId() int {
	return n.Id
}
