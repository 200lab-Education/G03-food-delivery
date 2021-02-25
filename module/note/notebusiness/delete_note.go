package notebusiness

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
	"errors"
)

//type NoteStore interface {
//	FindDataWithCondition(condition map[string]interface{}) (*notemodel.Note, error)
//	Delete(id int) error
//	ListDataWithCondition(cond map[string]interface{}, paging *common.Paging) ([]notemodel.Note, error)
//	// 10 more functions
//}

type DeleteNoteStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*notemodel.Note, error)
	Delete(ctx context.Context, id int) error
}

type deleteNoteBiz struct {
	store     DeleteNoteStore
	requester common.Requester
}

func NewDeleteNoteBiz(store DeleteNoteStore, requester common.Requester) *deleteNoteBiz {
	return &deleteNoteBiz{store: store, requester: requester}
}

func (biz *deleteNoteBiz) DeleteNote(ctx context.Context, noteId int) error {
	// Find note by id
	// If note note found: return error note not found
	// If old data has status is 0
	// 	=> error: note has been deleted
	// else
	// delete note
	note, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": noteId})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(notemodel.EntityName, err)
		}

		return common.ErrCannotDeleteEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return common.ErrCannotDeleteEntity(notemodel.EntityName, errors.New("note has been deleted before"))
	}

	if note.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(errors.New("you are not owner"))
	}

	if err := biz.store.Delete(ctx, note.Id); err != nil {
		return common.ErrCannotDeleteEntity(notemodel.EntityName, err)
	}

	return nil
}
