package cloudfiletag

import (
	"context"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileTagByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCloudFileTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileTagByIdLogic {
	return &GetCloudFileTagByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetCloudFileTagByIdLogic) GetCloudFileTagById(req *types.IDReq) (resp *types.CloudFileTagInfoResp, err error) {
	data, err := l.svcCtx.DB.CloudFileTag.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.CloudFileTagInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CloudFileTagInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Status: &data.Status,
			Name:   &data.Name,
			Remark: &data.Remark,
		},
	}, nil
}