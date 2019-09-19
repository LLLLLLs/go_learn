// Time        : 2019/09/09
// Description :

package profile

import (
	"context"
	"golearn/sundry/profile/errors"
	"golearn/sundry/profile/loader"
)

const STAT_VERSION = "statVersion"

func ContextWithVersion(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, STAT_VERSION, v)
}

func ContextWithTestVersion(ctx context.Context) context.Context {
	return context.WithValue(ctx, STAT_VERSION, loader.TEST_VERSION)
}

func getVersion(ctx context.Context) (string, error) {
	value := ctx.Value(STAT_VERSION)
	if value == nil {
		return "", errors.ErrContextMustWithVersion
	}
	return value.(string), nil
}
