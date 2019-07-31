package houyi

import (
	"fmt"
)

// ILayer defines all router handle interface includes single and group router.
type ILayer interface {
	Use(...HandlerFunc) ILayer
	NewLayer(string) *Layer
	Hit(string, HandlerFunc)
}

// Layer is used internally to configure router, a Layer is associated with
// a prefix and an array of handlers (middleware).
type Layer struct {
	handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

func (l *Layer) Use(handlers ...HandlerFunc) ILayer {
	l.handlers = append(l.handlers, handlers...)
	return l
}

func (l *Layer) NewLayer(name string) *Layer {
	return &Layer{
		handlers: l.handlers,
		basePath: l.combineLayer(name),
		engine:   l.engine,
	}
}

func (l *Layer) Hit(name string, handler HandlerFunc) {
	path := l.combineLayer(name)
	l.engine.lines[path] = append(l.handlers, handler)
	_, _ = fmt.Fprintf(DefaultWriter, "[HouYi] %-25s --> %s (%d handlers)\n", path, nameOfFunction(handler), len(l.engine.lines[path]))
}

var _ ILayer = &Layer{}

func (l *Layer) combineLayer(name string) string {
	if l.root {
		return name
	} else if name == "" {
		return l.basePath
	} else {
		return l.basePath + l.engine.separator + name
	}
}
