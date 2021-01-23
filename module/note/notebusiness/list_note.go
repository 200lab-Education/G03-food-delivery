package notebusiness

import (
	"demo/common"
	"demo/module/note/notemodel"
)

type ListNoteStore interface {
	ListDataWithCondition(cond map[string]interface{}, paging *common.Paging) ([]notemodel.Note, error)
}

type listNoteBiz struct {
	store ListNoteStore
}

func NewListNoteBiz(store ListNoteStore) *listNoteBiz {
	return &listNoteBiz{store: store}
}

func (biz *listNoteBiz) ListNote(paging *common.Paging) ([]notemodel.Note, error) {
	result, err := biz.store.ListDataWithCondition(nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(notemodel.EntityName, err)
	}

	return result, nil
}
