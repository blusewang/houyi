package houyi

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	y := New()
	y.
		Use(Recovery()).
		NewLayer("a").
		Use(func(context *Context) (err error) {
			log.Println("middle 1")
			err = context.Next()
			panic("panic test")
			return
		}).
		Use(func(context *Context) (err error) {
		err = context.Next()
			log.Println("middle 2")
			return
		}).
		NewLayer("b").
		Use(func(context *Context) (err error) {
			log.Println("middle 3")
			return
		}).
		Hit("", func(context *Context) (err error) {
			log.Println("hit")
			return
		}).Hit("c", func(context *Context) error {
		return nil
	})

	log.Println(y.Handle("a.b", []byte("data")))
}

func TestArr(t *testing.T) {
	var a = make(map[string]HandlersChain)
	a["sdf"] = HandlersChain{}
	log.Println(a["sdf"] == nil)
}
