package filetag

import (
	"context"
	"github.com/kebin6/simple-file-api/ent/filetag"
	"github.com/kebin6/simple-file-api/ent/predicate"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileTagListLogic {
	return &GetFileTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetFileTagListLogic) GetFileTagList(req *types.FileTagListReq) (resp *types.FileTagListResp, err error) {
	var predicates []predicate.FileTag
	if req.Name != nil {
		predicates = append(predicates, filetag.NameContains(*req.Name))
	}
	if req.Remark != nil {
		predicates = append(predicates, filetag.RemarkContains(*req.Remark))
	}
	data, err := l.svcCtx.DB.FileTag.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.FileTagListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.FileTagInfo{
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
