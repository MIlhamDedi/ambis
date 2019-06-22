package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"ambis/lib/config"
)

type proxyHandler struct{}

func (proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPaths := strings.Split(r.URL.Path, "/")
	rootPath := requestPaths[1]

	target := config.YuiEndpoint
	trimPath := true

	switch rootPath {
	case config.SinonRootPath:
		target = config.YuiEndpoint
	case config.AsunaRootPath:
		target = config.AsunaEndpoint
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

func NewHandler() *proxyHandler {
	return &proxyHandler{}
}