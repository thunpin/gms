package filters

import (
	"github.com/thunpin/gms"
	"github.com/thunpin/gms/logs"
)

func Log(tag string) gms.Action {
	return func(context *gms.Context, chain *gms.Chain) (interface{}, error) {
		result, err := chain.Next(context)
		logOut(tag, context, result, err)
		return result, err
	}
}

func logOut(tag string, context *gms.Context, result interface{}, err error) {
	header := logs.Header{tag, context.RequestId, context.UUID, context.Path}
	if err == nil {
		logs.Instance().Info(logs.Info{header, result})
	} else {
		logs.Instance().Error(logs.Error{header, err})
	}
}
