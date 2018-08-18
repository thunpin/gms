package xorm

import (
	"github.com/thunpin/gerrors"
)

type ToExecute func() (interface{}, gerrors.Errors)

func Transaction(uuid string, toExecute ToExecute) (interface{}, gerrors.Errors) {
	BeginSession(uuid)
	defer CloseSession(uuid)

	value, err := toExecute()
	if err != nil {
		Commit(uuid)
	} else {
		Rollback(uuid)
	}

	return value, err
}
