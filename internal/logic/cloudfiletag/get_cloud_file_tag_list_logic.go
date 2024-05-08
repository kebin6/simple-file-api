package cloudfiletag

import (
	"context"
	"github.com/kebin6/simple-file-api/ent/cloudfiletag"
	"github.com/kebin6/simple-file-api/ent/predicate"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCloudFileTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileTagListLogic {
	return &GetCloudFileTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetCloudFileTagListLogic) GetCloudFileTagList(req *types.CloudFileTagListReq) (resp *types.CloudFileTagListResp, err error) {
	var predicates []predicate.CloudFileTag
	if req.Name != nil {
		predicates = append(predicates, cloudfiletag.NameContains(*req.Name))
	}
	if req.Remark != nil {
		predicates = append(predicates, cloudfiletag.RemarkContains(*req.Remark))
	}
	data, err := l.svcCtx.DB.CloudFileTag.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.CloudFileTagListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.CloudFileTagInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Status: &v.Status,
				Name:   &v.Name,
				Remark: &v.Remark,
			})
	}

	return resp, nil
}
