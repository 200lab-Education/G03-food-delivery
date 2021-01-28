package notebusiness

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
)

type CreateNoteStore interface {
	Create(ctx context.Context, data *notemodel.CreateNote) error
}

type createNoteBusiness struct {
	store CreateNoteStore
}

func NewCreateNoteBusiness(store CreateNoteStore) *createNoteBusiness {
	return &createNoteBusiness{
		store: store,
	}
}

func (biz *createNoteBusiness) CreateNote(ctx context.Context, data *notemodel.CreateNote) error {

	data.Status = 1

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}
	return nil
}
