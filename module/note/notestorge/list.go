package notestorge

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
)

func (s *store) ListDataWithCondition(ctx context.Context, cond map[string]interface{}, paging *common.Paging) ([]notemodel.Note, error) {
	db := s.db.Table(notemodel.Note{}.TableName())

	db = db.Where("status <> 0")
	db = db.Where(cond)
	db = db.Order("id desc")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var data []notemodel.Note

	if err := db.Table(notemodel.Note{}.TableName()).
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
