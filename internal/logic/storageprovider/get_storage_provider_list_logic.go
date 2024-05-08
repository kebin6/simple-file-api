package storageprovider

import (
	"context"
	"github.com/kebin6/simple-file-api/ent/predicate"
	"github.com/kebin6/simple-file-api/ent/storageprovider"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorageProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStorageProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorageProviderListLogic {
	return &GetStorageProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetStorageProviderListLogic) GetStorageProviderList(req *types.StorageProviderListReq) (resp *types.StorageProviderListResp, err error) {
	var predicates []predicate.StorageProvider
	if req.Name != nil {
		predicates = append(predicates, storageprovider.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.StorageProvider.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.StorageProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.StorageProviderInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				State:       &v.State,
				Name:        &v.Name,
				Bucket:      &v.Bucket,
				SecretId:    &v.SecretID,
				SecretKey:   &v.SecretKey,
				Region:      &v.Region,
				IsDefault:   &v.IsDefault,
				Folder:      &v.Folder,
				Endpoint:    &v.Endpoint,
				PreviewHost: &v.PreviewHost,
			})
	}

	return resp, nil
}
