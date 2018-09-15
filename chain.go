package gms

import (
	"github.com/gin-gonic/gin"
)

func Route(
	path string,
	action Action,
	filters ...Filter) (string, func(*gin.Context)) {

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

type Action func(*Context) (interface{}, error)
type Filter func(*Context, *Chain) (interface{}, error)

type Context struct {
	*gin.Context
	RequestId string
	UUID      string
	Path      string
	Params    map[string]interface{}
}
type Chain struct {
	action Action
	filter Filter
	next   *Chain
}

func (chain *Chain) Next(context *Context) (interface{}, error) {
	if chain.next.action != nil {
		return chain.next.action(context)
	} else {
		return chain.next.filter(context, chain.next)
	}
}

func (chain *Chain) add(filter Filter) *Chain {
	if chain.next != nil {
		return chain.next.add(filter)
	} else {
		chain.next = &Chain{nil, filter, nil}
		return chain.next
	}
}

func newChain(action Action, filters ...Filter) *Chain {
	actionChain := &Chain{action, nil, nil}
	if len(filters) == 0 {
		return actionChain
	}

	root := &Chain{nil, filters[0], nil}
	chain := root
	for i := 1; i < len(filters); i++ {
		filter := filters[i]
		chain.add(filter)
		chain = chain.next
	}

	chain.next = actionChain
	return root
}

func execChain(chain *Chain, context *Context) (interface{}, error) {
	if chain.action != nil {
		return chain.action(context)
	} else {
		return chain.filter(context, chain)
	}
}
