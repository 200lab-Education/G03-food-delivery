package notestorge

import "demo/module/note/notemodel"

func (s *store) Delete(id int) error {
	db := s.db

	if err := db.Table(notemodel.Note{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}

	return nil
}
