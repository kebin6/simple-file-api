package file

import (
	"context"
	"github.com/kebin6/simple-file-api/ent"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/kebin6/simple-file-api/internal/utils/entx"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"os"
	"path"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePublicStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePublicStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePublicStatusLogic {
	return &ChangePublicStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ChangePublicStatusLogic) ChangePublicStatus(req *types.StatusCodeReq) (resp *types.BaseMsgResp, err error) {
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		file, err := tx.File.UpdateOneID(uuidx.ParseUUIDString(req.Id)).SetStatus(uint8(req.Status)).Save(l.ctx)

		if err != nil {
			return err
		}

		if req.Status == 1 {
			err = os.Rename(path.Join(l.svcCtx.Config.UploadConf.PrivateStorePath, file.Path),
				path.Join(l.svcCtx.Config.UploadConf.PublicStorePath, file.Path))
			if err != nil {
				logx.Errorw("fail to change the path of file", logx.Field("detail", err.Error()))
				return err
			}
		} else {
			err = os.Rename(path.Join(l.svcCtx.Config.UploadConf.PublicStorePath, file.Path),
				path.Join(l.svcCtx.Config.UploadConf.PrivateStorePath, file.Path))
			if err != nil {
				logx.Errorw("fail to change the path of file", logx.Field("detail", err.Error()))
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
	}, nil
}
