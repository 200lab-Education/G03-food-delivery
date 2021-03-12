package notestorge

import (
	"context"
	"demo/module/note/notemodel"
)

func (s *store) Delete(ctx context.Context, id int) error {
	db := s.db.Begin()

	if err := db.Table(notemodel.Note{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
