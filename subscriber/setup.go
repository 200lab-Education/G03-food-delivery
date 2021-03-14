package subscriber

import (
	"context"
	"demo/component/appctx"
)

func Setup(appCtx appctx.AppContext) {
	RunDeleteImageRecordAfterCreateNote(appCtx, context.Background())
	RunPushNotificationRecordAfterCreateNote(appCtx, context.Background())
}
