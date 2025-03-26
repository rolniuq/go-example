package main

import "strings"

// client will send request to proxy
// proxy will send request to real api which handle the request

type API interface {
	Request(string) string
}

type RealAPI struct{}

func (r *RealAPI) Request(req string) string {
	if req == "real" {
		return "I'm real. I say OK"
	}

	return "Not OK"
}

type ProxyAPI struct {
	api     API
	allowed []string
}

func NewProxyAPI(api API, allowed []string) *ProxyAPI {
	return &ProxyAPI{
		api:     api,
		allowed: allowed,
	}
}

func (p *ProxyAPI) Request(req string) string {
	if !strings.Contains(strings.Join(p.allowed, ","), req) {
		return "Access Denied"
	}

	return p.api.Request(req)
}

func main() {
	proxy := NewProxyAPI(&RealAPI{}, []string{"real", "fake"})

	println(proxy.Request("real"))
	println(proxy.Request("fake"))
}
