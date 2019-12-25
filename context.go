package houyi

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"unicode/utf8"
)

type Context struct {
	env      interface{}
	handlers HandlersChain
	index    int8
	uri      string
	data     []byte
	result   []byte
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
	return gob.NewDecoder(&b).Decode(o)
}

func (c *Context) BindJson(o interface{}) (err error) {
	return json.Unmarshal(c.data, o)
}

func (c *Context) BindProtoBuf(o proto.Message) (err error) {
	return proto.Unmarshal(c.data, o)
}

func (c *Context) String() (data string) {
	if utf8.Valid(c.data) {
		return string(c.data)
	}
	return
}

func (c *Context) GetUri() (data string) {
	return c.uri
}

func (c *Context) SetResult(raw []byte) {
	c.result = raw
}

func (c *Context) GetResult() []byte {
	return c.result
}

func (c *Context) GetEnv() interface{} {
	return c.env
}

func (c *Context) Next() (err error) {
	c.index++
	for c.index < int8(len(c.handlers)) {
		err = c.handlers[c.index](c)
		c.index++
	}
	return
}
