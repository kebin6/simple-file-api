package cloudfiletag

import (
	"context"
	"github.com/kebin6/simple-file-api/ent/cloudfiletag"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCloudFileTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCloudFileTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCloudFileTagLogic {
	return &DeleteCloudFileTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteCloudFileTagLogic) DeleteCloudFileTag(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	_, err = l.svcCtx.DB.CloudFileTag.Delete().Where(cloudfiletag.IDIn(req.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
