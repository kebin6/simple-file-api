package cloudfile

import (
	"context"
	"github.com/kebin6/simple-file-api/ent/cloudfile"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCloudFileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileByIdLogic {
	return &GetCloudFileByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetCloudFileByIdLogic) GetCloudFileById(req *types.UUIDReq) (resp *types.CloudFileInfoResp, err error) {
	data, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.IDEQ(uuidx.ParseUUIDString(req.Id))).WithStorageProviders().
		First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID.String()),
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			State:      &data.State,
			Name:       &data.Name,
			Url:        &data.URL,
			Size:       &data.Size,
			FileType:   &data.FileType,
			UserId:     &data.UserID,
			ProviderId: &data.Edges.StorageProviders.ID,
		},
	}, nil
}
