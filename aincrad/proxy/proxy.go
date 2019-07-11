package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"ambis/lib/base"
)

type proxyHandler struct {
	Base *base.Base
}

func (h proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	appConfig := h.Base.Config
	requestPaths := strings.Split(r.URL.Path, "/")
	rootPath := requestPaths[1]

	target := appConfig.YuiEndpoint
	trimPath := true

	switch rootPath {
	case appConfig.SinonRootPath:
		target = appConfig.YuiEndpoint
	case appConfig.AsunaRootPath:
		target = appConfig.AsunaEndpoint
	default: // Yui serves on '/'
		trimPath = false
	}

	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update Request Object
	// Request Path
	if trimPath {
		r.URL.Path = "/"
		if len(requestPaths) > 2 {
			r.URL.Path += strings.Join(requestPaths[2:], "/")
		}
	}
	proxy.ServeHTTP(w, r)
}

func New(b *base.Base) *proxyHandler {
	return &proxyHandler{Base: b}
}
