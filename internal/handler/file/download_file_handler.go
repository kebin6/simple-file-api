package file

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kebin6/simple-file-api/internal/logic/file"
	"github.com/kebin6/simple-file-api/internal/svc"
	"github.com/kebin6/simple-file-api/internal/types"
)

// swagger:route get /file/download/{id} file DownloadFile
//
// Download file | 下载文件
//
// Download file | 下载文件
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDPathReq
//

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDPathReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownloadFileLogic(r.Context(), svcCtx)
		filePath, err := l.DownloadFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			body, err := os.ReadFile(filePath)
			if err != nil {
				httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
				return
			}
			w.Header().Set("Accept-Encoding", "identity;q=1, *;q=0")
			w.Write(body)
		}
	}
}
