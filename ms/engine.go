package ms

import "github.com/gin-gonic/gin"

type Engine struct {
	Engine *gin.Engine
}

func NewEngine(engine *gin.Engine) *Engine {
	return &Engine{engine}
}

func (engine Engine) GET(path string, action Action) *Engine {
	engine.Engine.GET(path, internalAction(path, "GET", action))
	return &engine
}

func (engine Engine) POST(path string, action Action) *Engine {
	engine.Engine.POST(path, internalAction(path, "POST", action))
	return &engine
}

func (engine Engine) PUT(path string, action Action) *Engine {
	engine.Engine.PUT(path, internalAction(path, "PUT", action))
	return &engine
}

func (engine Engine) DELETE(path string, action Action) *Engine {
	engine.Engine.DELETE(path, internalAction(path, "DELETE", action))
	return &engine
}

func internalAction(
	tagName string,
	actionName string,
	action Action) func(*gin.Context) {

	return func(context *gin.Context) {
		Response(context).
			TagName(tagName).
			ActionName(actionName).
			Action(action).
			Run()
	}
}
