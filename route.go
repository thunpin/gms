package gms

import (
	"github.com/gin-gonic/gin"
)

func Route(
	path string,
	action Action,
	filters ...Action) (string, func(*gin.Context)) {

	chain := newChain(action, filters...)

	return path, func(ginContext *gin.Context) {
		requestId := ginContext.GetHeader(RequestIdHeader)
		uuid := NewUUID()

		context := Context{
			ginContext,
			requestId,
			uuid,
			path,
			make(map[string]interface{}),
		}

		execChain(chain, &context)
	}
}

type Action func(*Context, *Chain) (interface{}, error)

type Context struct {
	*gin.Context
	RequestId string
	UUID      string
	Path      string
	Params    map[string]interface{}
}
type Chain struct {
	action Action
	next   *Chain
}

func (chain Chain) Next(context *Context) (interface{}, error) {
	return chain.next.action(context, chain.next)
}

func (chain Chain) add(action Action) *Chain {
	if chain.next != nil {
		return chain.next.add(action)
	} else {
		chain.next = &Chain{action, nil}
		return chain.next
	}
}

func newChain(action Action, filters ...Action) *Chain {
	if len(filters) == 0 {
		return &Chain{action, nil}
	}

	root := &Chain{filters[0], nil}
	chain := root
	for i := 1; i < len(filters); i++ {
		filter := filters[i]
		chain.add(filter)
		chain = chain.next
	}
	chain.add(action)
	return root
}

func execChain(chain *Chain, context *Context) (interface{}, error) {
	return chain.action(context, chain)
}
