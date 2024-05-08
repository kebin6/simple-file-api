package cloudfile

import (
	"context"
	"github.com/kebin6/simple-file-api/ent"
	"github.com/kebin6/simple-file-api/ent/cloudfile"
	"github.com/kebin6/simple-file-api/ent/cloudfiletag"
	"github.com/kebin6/simple-file-api/ent/predicate"
	"github.com/kebin6/simple-file-api/ent/storageprovider"
	"github.com/kebin6/simple-file-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCloudFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileListLogic {
	return &GetCloudFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetCloudFileListLogic) GetCloudFileList(req *types.CloudFileListReq) (resp *types.CloudFileListResp, err error) {
	var predicates []predicate.CloudFile
	if req.Name != nil {
		predicates = append(predicates, cloudfile.NameContains(*req.Name))
	}
	if req.ProviderId != nil {
		predicates = append(predicates, cloudfile.HasStorageProvidersWith(storageprovider.IDEQ(*req.ProviderId)))
	}
	if req.TagIds != nil {
		predicates = append(predicates, cloudfile.HasTagsWith(cloudfiletag.IDIn(req.TagIds...)))
	}
	if req.FileType != nil && *req.FileType != 0 {
		predicates = append(predicates, cloudfile.FileTypeEQ(*req.FileType))
	}
	data, err := l.svcCtx.DB.CloudFile.Query().Where(predicates...).WithStorageProviders().WithTags().
		Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.CloudFileListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		previewURL := v.Edges.StorageProviders.PreviewHost + v.Path
		resp.Data.Data = append(resp.Data.Data,
			types.CloudFileInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        pointy.GetPointer(v.ID.String()),
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				State:      &v.State,
				Name:       &v.Name,
				Url:        &previewURL,
				Size:       &v.Size,
				FileType:   &v.FileType,
				UserId:     &v.UserID,
				ProviderId: &v.Edges.StorageProviders.ID,
				TagIds:     l.getFileTagIds(v.Edges.Tags),
			})
	}

	return resp, nil
}

func (l *GetCloudFileListLogic) getFileTagIds(tags []*ent.CloudFileTag) (result []uint64) {
	for _, v := range tags {
		result = append(result, v.ID)
	}
	return result
}
