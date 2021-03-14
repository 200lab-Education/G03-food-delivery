package subscriber

import (
	"context"
	"demo/common"
	"demo/component/appctx"
	"demo/module/upload/uploadstorage"
	"demo/pubsub"
	"log"
)

type HasImageIds interface {
	GetId() int
	GetImageIds() []int
}

func RunDeleteImageRecordAfterCreateNote(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicNoteCreated)

	go func() {
		for {
			msg := <-c
			hasImgData := msg.Data().(HasImageIds)

			uploadstorage.NewSQLStore(appCtx.GetMainDBConnection()).DeleteImages(ctx, hasImgData.GetImageIds())

			//if data, ok := msg.Data().(HasImageIds); ok {
			//
			//
			//	uploadstorage.NewSQLStore(appCtx.GetMainDBConnection()).DeleteImages(ctx, data.GetImageIds())
			//}
		}
	}()
}

func RunPushNotificationRecordAfterCreateNote(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicNoteCreated)

	go func() {
		for {
			msg := <-c
			hasImgData := msg.Data().(HasImageIds)

			log.Println("push notification after note created with id:", hasImgData.GetId())

			//uploadstorage.NewSQLStore(appCtx.GetMainDBConnection()).DeleteImages(ctx, []int{newNote.CoverImgId})

			//if data, ok := msg.Data().(HasImageIds); ok {
			//
			//
			//	uploadstorage.NewSQLStore(appCtx.GetMainDBConnection()).DeleteImages(ctx, data.GetImageIds())
			//}
		}
	}()
}

func DeleteImageRecordAfterCreateNote(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete images records after create note",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			if data, ok := message.Data().(HasImageIds); ok {
				return uploadstorage.NewSQLStore(appCtx.GetMainDBConnection()).DeleteImages(ctx, data.GetImageIds())
			}

			return nil
		},
	}
}

func SendEmailAfterCreateNote(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Send email after create note",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			hasImgData := message.Data().(HasImageIds)
			log.Println("sending email for note id:", hasImgData.GetId())
			return nil
		},
	}
}
