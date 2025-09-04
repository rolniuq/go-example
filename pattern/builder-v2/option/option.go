package option

import "sony/internal"

type Option interface {
	Apply(*internal.DialSettings)
}

type withToken string

func (w withToken) Apply(o *internal.DialSettings) {
	o.Token = string(w)
}

func WithToken(token string) Option {
	return withToken(token)
}

type withCredentials string

func (w withCredentials) Apply(o *internal.DialSettings) {
	o.Credentials = string(w)
}

func WithCredentials(credentials string) Option {
	return withCredentials(credentials)
}
