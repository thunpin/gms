package filters

import (
	"github.com/thunpin/gms"
	"github.com/thunpin/gms/logs"
)

func LogAll(tag string) gms.Action {
	return func(context *gms.Context, chain *gms.Chain) (interface{}, error) {
		//logIn(tag, context)
		result, err := chain.Next(context)
		logOut(tag, context, result, err)
		return result, err
	}
}

func LogIn(tag string) gms.Action {
	return func(context *gms.Context, chain *gms.Chain) (interface{}, error) {
		logIn(tag, context)
		return chain.Next(context)
	}
}

func LogOut(tag string) gms.Action {
	return func(context *gms.Context, chain *gms.Chain) (interface{}, error) {
		result, err := chain.Next(context)
		logOut(tag, context, result, err)
		return result, err
	}
}

func logIn(tag string, context *gms.Context) {
	data, _ := context.GetRawData()
	header := logs.Header{tag, context.RequestId, context.UUID, context.Path}
	logs.Instance().Info(logs.Info{header, string(data)})
}

func logOut(tag string, context *gms.Context, result interface{}, err error) {
	header := logs.Header{tag, context.RequestId, context.UUID, context.Path}
	if err == nil {
		logs.Instance().Info(logs.Info{header, result})
	} else {
		logs.Instance().Error(logs.Error{header, err})
	}
}
