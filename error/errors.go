// Time        : 2019/01/04
// Description :

package error

import "github.com/pkg/errors"

func first() error {
	return errors.New("first")
}

func second() error {
	return errors.Wrap(first(), "second")
}

func third() error {
	return errors.Wrap(second(), "third")
}
