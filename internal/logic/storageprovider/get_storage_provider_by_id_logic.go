package storageprovider

import (
	"context"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorageProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStorageProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorageProviderByIdLogic {
	return &GetStorageProviderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetStorageProviderByIdLogic) GetStorageProviderById(req *types.IDReq) (resp *types.StorageProviderInfoResp, err error) {
	data, err := l.svcCtx.DB.StorageProvider.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.StorageProviderInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StorageProviderInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			State:       &data.State,
			Name:        &data.Name,
			Bucket:      &data.Bucket,
			SecretId:    &data.SecretID,
			SecretKey:   &data.SecretKey,
			Region:      &data.Region,
			IsDefault:   &data.IsDefault,
			Folder:      &data.Folder,
			Endpoint:    &data.Endpoint,
			PreviewHost: &data.PreviewHost,
		},
	}, nil
}
