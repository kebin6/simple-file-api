package filetag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/simple-file-api/internal/logic/filetag"
	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"
)

// swagger:route post /file_tag/create filetag CreateFileTag
//
// Create file tag information | 创建文件标签
//
// Create file tag information | 创建文件标签
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: FileTagInfo
//
// Responses:
//  200: BaseMsgResp

func CreateFileTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileTagInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := filetag.NewCreateFileTagLogic(r.Context(), svcCtx)
		resp, err := l.CreateFileTag(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
