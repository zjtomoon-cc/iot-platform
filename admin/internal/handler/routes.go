// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"gitee/getcharzp/iot-platform/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/device/list",
				Handler: DeviceListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/product/list",
				Handler: ProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/product/create",
				Handler: ProductCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/product/modify",
				Handler: ProductModifyHandler(serverCtx),
			},
		},
	)
}
