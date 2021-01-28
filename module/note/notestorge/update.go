package notestorge

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
)

func (s *store) Update(ctx context.Context, data *notemodel.UpdateNote) error {
	db := s.db.Begin()
	db = db.Table(data.TableName())
	db = db.Where("id = ?", data.Id)

	if err := db.Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
