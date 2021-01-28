package notestorge

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
)

func (s *store) Create(ctx context.Context, data *notemodel.CreateNote) error {
	db := s.db.Begin()
	db = db.Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
