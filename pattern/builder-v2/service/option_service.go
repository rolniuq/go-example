package service

import "sony/option"

type OptionService struct {
}

func NewOptionService(opts ...option.Option) *OptionService {
	return &OptionService{}
}

func (s *OptionService) Build() string {
	return ""
}

func Test() {
	s := NewOptionService(option.WithCredentials("test"), option.WithToken("test"))
	s.Build()
}
