package cloudfiletag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/simple-file-api/internal/logic/cloudfiletag"
	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"
)

// swagger:route post /cloud_file_tag/create cloudfiletag CreateCloudFileTag
//
// Create cloud file tag information | 创建云文件标签
//
// Create cloud file tag information | 创建云文件标签
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CloudFileTagInfo
//
// Responses:
//  200: BaseMsgResp

func CreateCloudFileTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CloudFileTagInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfiletag.NewCreateCloudFileTagLogic(r.Context(), svcCtx)
		resp, err := l.CreateCloudFileTag(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
