// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "github.com/kebin6/simple-file-api/internal/handler/base"
	cloudfile "github.com/kebin6/simple-file-api/internal/handler/cloudfile"
	cloudfiletag "github.com/kebin6/simple-file-api/internal/handler/cloudfiletag"
	file "github.com/kebin6/simple-file-api/internal/handler/file"
	filetag "github.com/kebin6/simple-file-api/internal/handler/filetag"
	storageprovider "github.com/kebin6/simple-file-api/internal/handler/storageprovider"
	"github.com/kebin6/simple-file-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/list",
					Handler: file.FileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/update",
					Handler: file.UpdateFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/delete",
					Handler: file.DeleteFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/status",
					Handler: file.ChangePublicStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/file/download/:id",
					Handler: file.DownloadFileHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: file.UploadHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file_tag/create",
					Handler: filetag.CreateFileTagHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file_tag/update",
					Handler: filetag.UpdateFileTagHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file_tag/delete",
					Handler: filetag.DeleteFileTagHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file_tag/list",
					Handler: filetag.GetFileTagListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file_tag",
					Handler: filetag.GetFileTagByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file/create",
					Handler: cloudfile.CreateCloudFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file/update",
					Handler: cloudfile.UpdateCloudFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file/delete",
					Handler: cloudfile.DeleteCloudFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file/list",
					Handler: cloudfile.GetCloudFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file",
					Handler: cloudfile.GetCloudFileByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file/upload",
					Handler: cloudfile.UploadHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/storage_provider/create",
					Handler: storageprovider.CreateStorageProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/storage_provider/update",
					Handler: storageprovider.UpdateStorageProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/storage_provider/delete",
					Handler: storageprovider.DeleteStorageProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/storage_provider/list",
					Handler: storageprovider.GetStorageProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/storage_provider",
					Handler: storageprovider.GetStorageProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file_tag/create",
					Handler: cloudfiletag.CreateCloudFileTagHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file_tag/update",
					Handler: cloudfiletag.UpdateCloudFileTagHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file_tag/delete",
					Handler: cloudfiletag.DeleteCloudFileTagHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file_tag/list",
					Handler: cloudfiletag.GetCloudFileTagListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cloud_file_tag",
					Handler: cloudfiletag.GetCloudFileTagByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}