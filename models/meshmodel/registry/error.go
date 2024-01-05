package registry

import (
	"github.com/khulnasoft/meshkit/errors"
)

var (
	ErrUnknownHostCode = "11097"
)

func ErrUnknownHost(err error) error {
	return errors.New(ErrUnknownHostCode, errors.Alert, []string{"host is not supported"}, []string{err.Error()}, []string{"The component's host is not supported by the version of server you are running"}, []string{"Try upgrading to latest available version"})
}
