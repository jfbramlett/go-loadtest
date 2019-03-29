package testwrapper

import "context"

type Test interface {
	Run(ctx context.Context) error
}
