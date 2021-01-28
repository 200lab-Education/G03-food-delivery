package notemodel

type UpdateNote struct {
	Id      int     `json:"id" form:"id" gorm:"column:id"`
	Title   *string `json:"title" form:"title" gorm:"column:title"`
	Content *string `json:"content" form:"content" gorm:"column:content"`
}

func (UpdateNote) TableName() string {
	return Note{}.TableName()
}
