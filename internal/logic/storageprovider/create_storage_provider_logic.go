package storageprovider

import (
	"context"
	"github.com/kebin6/simple-file-api/internal/utils/cloud"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStorageProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStorageProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStorageProviderLogic {
	return &CreateStorageProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateStorageProviderLogic) CreateStorageProvider(req *types.StorageProviderInfo) (resp *types.BaseMsgResp, err error) {
	_, err = l.svcCtx.DB.StorageProvider.Create().
		SetNotNilState(req.State).
		SetNotNilName(req.Name).
		SetNotNilBucket(req.Bucket).
		SetNotNilSecretID(req.SecretId).
		SetNotNilSecretKey(req.SecretKey).
		SetNotNilRegion(req.Region).
		SetNotNilIsDefault(req.IsDefault).
		SetNotNilFolder(req.Folder).
		SetNotNilEndpoint(req.Endpoint).
		SetNotNilPreviewHost(req.PreviewHost).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	l.svcCtx.CloudStorage = cloud.NewCloudServiceGroup(l.svcCtx.DB)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
