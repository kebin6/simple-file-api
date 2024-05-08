package cloudfile

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kebin6/simple-file-api/ent/cloudfile"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/errorx"
	"strings"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCloudFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCloudFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCloudFileLogic {
	return &DeleteCloudFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteCloudFileLogic) DeleteCloudFile(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.UploadConf.DeleteFileWithCloud {
		data, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).
			WithStorageProviders().All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		for _, v := range data {
			if client, ok := l.svcCtx.CloudStorage.CloudStorage[v.Edges.StorageProviders.Name]; ok {
				keyData := strings.Split(v.URL, *l.svcCtx.CloudStorage.CloudStorage[v.Edges.StorageProviders.Name].Config.Endpoint)
				if len(keyData) != 2 {
					logx.Errorw("failed to find the key of the cloud file", logx.Field("data", req))
					return nil, errorx.NewCodeInternalError(i18n.Failed)
				}
				_, err = client.DeleteObject(&s3.DeleteObjectInput{
					Bucket: aws.String(l.svcCtx.CloudStorage.ProviderData[v.Edges.StorageProviders.Name].Bucket),
					Key:    aws.String(keyData[1]),
				})
				if err != nil {
					logx.Errorw("failed to delete the cloud file", logx.Field("detail", err), logx.Field("data", req))
					return nil, errorx.NewCodeInternalError(i18n.Failed)
				}
			}
		}
	}

	_, err = l.svcCtx.DB.CloudFile.Delete().Where(cloudfile.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
