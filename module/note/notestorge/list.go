package notestorge

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
)

func (s *store) ListDataWithCondition(ctx context.Context, cond map[string]interface{},
	paging *common.Paging, moreDatas ...string) ([]notemodel.Note, error) {
	db := s.db.Table(notemodel.Note{}.TableName())

	db = db.Where("status <> 0")
	db = db.Where(cond)
	db = db.Order("id desc")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var data []notemodel.Note

	db = db.Table(notemodel.Note{}.TableName()).Limit(paging.Limit)

	if paging.FakeCursor != "" {
		uid, err := common.FromBase58(paging.FakeCursor)

		if err != nil {
			return nil, err
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	for i := range moreDatas {
		db = db.Preload(moreDatas[i])
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
