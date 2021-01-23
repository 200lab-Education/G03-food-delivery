package notebusiness

import (
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
	FindDataWithCondition(condition map[string]interface{}) (*notemodel.Note, error)
	Delete(id int) error
}

type deleteNoteBiz struct {
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(noteId int) error {
	// Find note by id
	// If note note found: return error note not found
	// If old data has status is 0
	// 	=> error: note has been deleted
	// else
	// delete note
	note, err := biz.store.FindDataWithCondition(map[string]interface{}{"id": noteId})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(notemodel.EntityName, err)
		}

		return common.ErrCannotDeleteEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return errors.New("note has been deleted before")
	}

	if err := biz.store.Delete(note.Id); err != nil {
		return err
	}

	return nil
}
