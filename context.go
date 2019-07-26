package houyi

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type Context struct {
	handlers HandlersChain
	index    int8
	path     string
	data     []byte
	engine   *Engine
}

func (c *Context) RawData() []byte {
	return c.data
}

func (c *Context) Bind(o interface{}) (err error) {
	var b bytes.Buffer
	if _, err = b.Write(c.data); err != nil {
		return
	}
	err = gob.NewDecoder(&b).Decode(o)
	return
}

func (c *Context) BindJson(o interface{}) (err error) {
	err = json.Unmarshal(c.data, o)
	return
}

func (c *Context) Next() (err error) {
	c.index++
	for c.index < int8(len(c.handlers)) {
		err = c.handlers[c.index](c)
		if err != nil {
			c.index = int8(len(c.handlers))
		}else{
			c.index++
		}
	}
	return
}