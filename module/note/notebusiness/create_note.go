package notebusiness

import (
	"context"
	"demo/common"
	"demo/module/note/notemodel"
	"demo/pubsub"
)

type ImageStore interface {
	ListImages(
		context context.Context,
		ids []int,
		moreKeys ...string,
	) ([]common.Image, error)

	DeleteImages(ctx context.Context, ids []int) error
}

type CreateNoteStore interface {
	Create(ctx context.Context, data *notemodel.CreateNote) error
}

type createNoteBusiness struct {
	imgStore ImageStore
	store    CreateNoteStore
	pb       pubsub.Pubsub
}

func NewCreateNoteBusiness(store CreateNoteStore, imgStore ImageStore, pb pubsub.Pubsub) *createNoteBusiness {
	return &createNoteBusiness{
		store:    store,
		imgStore: imgStore,
		pb:       pb,
	}
}

func (biz *createNoteBusiness) CreateNote(ctx context.Context, data *notemodel.CreateNote) error {
	data.Status = 1

	imgs, err := biz.imgStore.ListImages(ctx, []int{data.CoverImgId})

	if err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	if len(imgs) == 0 {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	data.Cover = &imgs[0]

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	// Side effect, we need to off-load
	go func() {
		biz.pb.Publish(ctx, common.TopicNoteCreated, pubsub.NewMessage(data))
	}()

	//go func() {
	//	err := biz.imgStore.DeleteImages(ctx, []int{data.CoverImgId})
	//
	//	if err != nil {
	//		_ = biz.imgStore.DeleteImages(ctx, []int{data.CoverImgId})
	//	}
	//}()

	return nil
}
