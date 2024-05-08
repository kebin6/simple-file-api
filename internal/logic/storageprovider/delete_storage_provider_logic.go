package storageprovider

import (
	"context"
	"github.com/kebin6/simple-file-api/ent/cloudfile"
	"github.com/kebin6/simple-file-api/ent/storageprovider"
	"github.com/kebin6/simple-file-api/internal/utils/cloud"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStorageProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStorageProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStorageProviderLogic {
	return &DeleteStorageProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteStorageProviderLogic) DeleteStorageProvider(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	check, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.HasStorageProvidersWith(storageprovider.IDIn(req.Ids...))).
		Count(l.ctx)

	if err != nil {
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
	}

	if check != 0 {
		return nil, errorx.NewCodeInvalidArgumentError(l.svcCtx.Trans.Trans(l.ctx, "storage_provider.hasFileError"))
	}

	_, err = l.svcCtx.DB.StorageProvider.Delete().Where(storageprovider.IDIn(req.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	l.svcCtx.CloudStorage = cloud.NewCloudServiceGroup(l.svcCtx.DB)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
