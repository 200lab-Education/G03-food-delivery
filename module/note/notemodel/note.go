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
	UserId          int                `json:"-" gorm:"column:user_id;"`
	FakeUserId      *common.UID        `json:"user_id" gorm:"-"`
	Title           string             `json:"title" form:"title" gorm:"column:title"`
	Content         string             `json:"content" form:"content" gorm:"column:content"`
	Cover           *common.Image      `json:"cover" gorm:"column:cover;"`
	Photos          common.Images      `json:"photos" gorm:"column:photos;"`
	User            *common.SimpleUser `json:"user" gorm:"foreignKey:UserId;"`
}

func (Note) TableName() string {
	return "notes"
}

func (n *Note) Mask() {
	n.GenUID(common.DbTypeNote)

	uid := common.NewUID(uint32(n.UserId), common.DbTypeUser, 1)
	n.FakeUserId = &uid

	if n.User != nil {
		n.User.GenUID(common.DbTypeUser)
	}
}
