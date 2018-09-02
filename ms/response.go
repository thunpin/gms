package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/thunpin/gerrors"
	"github.com/thunpin/gms/logs"
	"github.com/thunpin/gms/ms/jwt"
)

type Action func(*Context) (interface{}, gerrors.Errors)
type SecurityAction func(string, *Context) bool

type Context struct {
	RequestId      string
	UUID           string
	Context        *gin.Context
	Params         map[string]interface{}
	logHeader      logs.Header
	securityAction SecurityAction
	action         Action
}

func Response(context *gin.Context) *Context {
	uuid := NewUUID()
	requestId := context.GetHeader(RequestIdHeader)
	logHeader := logs.Header{"", requestId, uuid, ""}
	return &Context{
		RequestId: requestId,
		UUID:      uuid,
		Context:   context,
		Params:    make(map[string]interface{}),
		logHeader: logHeader,
	}
}

func (context Context) TagName(name string) *Context {
	context.logHeader.Tag = name
	return &context
}

func (context Context) ActionName(name string) *Context {
	context.logHeader.Action = name
	return &context
}

func (context Context) Info(info interface{}) *Context {
	logs.Instance().Info(logs.Info{context.logHeader, info})
	return &context
}

func (context Context) SecurityAction(securityAction SecurityAction) *Context {
	context.securityAction = securityAction
	return &context
}

func (context Context) Action(action Action) *Context {
	context.action = action
	return &context
}

func (context Context) Run() {
	err := executeSecurityAction(&context)

	var value interface{}
	if err == nil && context.action != nil {
		value, err = context.action(&context)
		if err != nil {
			logs.Instance().Error(logs.Error{context.logHeader, err})
		}
	}

	ToJSON(context.Context, value, err)
}

func executeSecurityAction(context *Context) error {
	if context.securityAction != nil {
		token, err := jwt.TokenFromHeader(context.Context)
		if err != nil {
			return gerrors.Forbidden()
		}

		if !context.securityAction(token, context) {
			return gerrors.Forbidden()
		}
	}
	return nil
}
