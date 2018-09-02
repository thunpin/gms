package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/thunpin/gerrors"
	"github.com/thunpin/gms/logs"
)

type Action func(*Context) (interface{}, gerrors.Errors)

type Context struct {
	RequestId string
	UUID      string
	Context   *gin.Context
	logHeader logs.Header
	action    Action
}

func Response(context *gin.Context) *Context {
	uuid := NewUUID()
	requestId := context.GetHeader(RequestIdHeader)
	logHeader := logs.Header{"", requestId, uuid, ""}
	return &Context{
		RequestId: requestId,
		UUID:      uuid,
		Context:   context,
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

func (context Context) Action(action Action) *Context {
	context.action = action
	return &context
}

func (context Context) Run() {
	if context.action != nil {
		value, err := context.action(&context)
		if err != nil {
			logs.Instance().Error(logs.Error{context.logHeader, err})
		}

		ToJSON(context.Context, value, err)
	}
}
