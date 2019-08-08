package houyi

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var DefaultWriter io.Writer = os.Stdout

var DefaultErrorWriter io.Writer = os.Stderr

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context) error

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// Last returns the last handler in the chain. ie. the last handler is the main own.
func (c HandlersChain) Last() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}

type Engine struct {
	Layer
	lines     map[string]HandlersChain
	separator string
	pool      sync.Pool
}

var _ ILayer = &Engine{}

func New() *Engine {
	e := &Engine{
		Layer:     Layer{root: true},
		separator: ".",
		lines:     make(map[string]HandlersChain),
	}
	e.engine = e
	e.pool.New = func() interface{} {
		return &Context{engine: e}
	}
	return e
}

func (e *Engine) Handle(path string, data []byte, env interface{}) (result []byte, err error) {
	if e.lines[path] != nil {
		c := e.pool.Get().(*Context)
		c.env = env
		c.path = path
		c.data = data
		c.index = -1
		c.handlers = e.lines[path]

		err = c.Next()
		result = c.GetResult()

		e.pool.Put(c)
	} else {
		log.Println(e.lines)
		err = fmt.Errorf("[%v]没有可命中的服务", path)
	}
	return
}
