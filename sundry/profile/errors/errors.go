// Time        : 2019/09/10
// Description :

package errors

import "errors"

var (
	ErrNoVersion              = errors.New("can't find version")
	ErrNoTable                = errors.New("can't find table")
	ErrVersionConfig          = errors.New("wrong version table")
	ErrNoConfig               = errors.New("no config")
	ErrContextMustWithVersion = errors.New("context must with version")
)
