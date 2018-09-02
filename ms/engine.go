package ms

import "github.com/gin-gonic/gin"

type Engine struct {
	Engine *gin.Engine
}

func NewEngine(engine *gin.Engine) *Engine {
	return &Engine{engine}
}

func (engine Engine) GET(path string, action Action) *Engine {
	engine.SecurityGET(path, nil, action)
	return &engine
}

func (engine Engine) POST(path string, action Action) *Engine {
	engine.SecurityPOST(path, nil, action)
	return &engine
}

func (engine Engine) PUT(path string, action Action) *Engine {
	engine.SecurityPUT(path, nil, action)
	return &engine
}

func (engine Engine) DELETE(path string, action Action) *Engine {
	engine.SecurityDELETE(path, nil, action)
	return &engine
}

func (engine Engine) SecurityGET(
	path string,
	securityAction SecurityAction,
	action Action) *Engine {

	engine.Engine.GET(path, internalAction(path, "GET", securityAction, action))
	return &engine
}

func (engine Engine) SecurityPOST(
	path string,
	securityAction SecurityAction,
	action Action) *Engine {

	engine.Engine.POST(path, internalAction(path, "POST", securityAction, action))
	return &engine
}

func (engine Engine) SecurityPUT(
	path string,
	securityAction SecurityAction,
	action Action) *Engine {

	engine.Engine.PUT(path, internalAction(path, "PUT", securityAction, action))
	return &engine
}

func (engine Engine) SecurityDELETE(
	path string,
	securityAction SecurityAction,
	action Action) *Engine {

	engine.Engine.DELETE(path, internalAction(path, "DELETE", securityAction, action))
	return &engine
}

func internalAction(
	tagName string,
	actionName string,
	securityAction SecurityAction,
	action Action) func(*gin.Context) {

	return func(context *gin.Context) {
		Response(context).
			TagName(tagName).
			ActionName(actionName).
			SecurityAction(securityAction).
			Action(action).
			Run()
	}
}
