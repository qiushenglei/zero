package handler

import (
	"net/http"

	"github.com/qiushenglei/gozero/http/zerohttp/internal/logic"
	"github.com/qiushenglei/gozero/http/zerohttp/internal/svc"
	"github.com/qiushenglei/gozero/http/zerohttp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ZerohttpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZerohttpLogic(r.Context(), svcCtx)
		resp, err := l.Zerohttp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
