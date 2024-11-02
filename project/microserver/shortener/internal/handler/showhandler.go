package handler

import (
	"net/http"

	"github.com/guojinghu/shortenter/internal/logic"
	"github.com/guojinghu/shortenter/internal/svc"
	"github.com/guojinghu/shortenter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析请求参数
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 执行业务逻辑
		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			// 返回重定向的响应(2.返回重定向响应)
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
		}
	}
}
