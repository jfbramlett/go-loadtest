package testwrapper

import (
	"context"
)

type httpTest struct {
}

func (h *httpTest) Run(ctx context.Context) error {
	return nil
}


func NewHttpTest() Test {
	return &httpTest{}
}
