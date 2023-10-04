package dynamicproxy

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

/* This is a Caddy module to enable per-request dynamic proxying.
The upstream host is determined by a header, x-upstream-host, and the
rest of the request is proxied to that host, including the path and query
string. The x-upstream-host header is stripped before proxying*/

func init() {
	caddy.RegisterModule(Proxy{})
	httpcaddyfile.RegisterHandlerDirective("dynamicproxy", parseCaddyfileHandlerDirective)
}

// Proxy is a Caddy module that enables dynamic proxying
type Proxy struct {
	UpstreamHostHeader string `json:"upstream_host_header,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Proxy) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.dynamicproxy",
		New: func() caddy.Module { return new(Proxy) },
	}
}

func (p *Proxy) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.UpstreamHostHeader = d.Val()
		}
	}
	return nil
}

func parseCaddyfileHandlerDirective(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var p Proxy
	err := p.UnmarshalCaddyfile(h.Dispenser)
	return p, err
}

func (m Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	upstreamHost := r.Header.Get(m.UpstreamHostHeader)
	if upstreamHost == "" {
		return next.ServeHTTP(w, r)
	}
	r.Header.Del(m.UpstreamHostHeader)

	caddyhttp.SetVar(r.Context(), "upstream_host", upstreamHost)

	return next.ServeHTTP(w, r)
}
