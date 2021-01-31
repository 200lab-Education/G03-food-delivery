package notebusiness

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
)

type UpdateNoteStore interface {
	Update(ctx context.Context, data *notemodel.UpdateNote) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*notemodel.Note, error)
}

type updateNoteBusiness struct {
	store UpdateNoteStore
}

func NewUpdateNoteBusiness(store UpdateNoteStore) *updateNoteBusiness {
	return &updateNoteBusiness{store: store}
}

func (biz *updateNoteBusiness) UpdateNote(ctx context.Context, data *notemodel.UpdateNote) error {
	note, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"id": data.Id,
	})

	if err != nil {
		return common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return notemodel.ErrNoteDeactive
	}

	if err = biz.store.Update(ctx, data); err != nil {
		return common.ErrCannotUpdateEntity(notemodel.EntityName, err)
	}

	return nil
}
