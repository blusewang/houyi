package houyi

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	y := New()
	y.
		NewLayer("a").
		Use(func(context *Context) (err error) {
			log.Println("middle 1")
			return
		}).
		Use(func(context *Context) (err error) {
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
		})

	log.Println(y.Handle("a.b", []byte("data")))
}

func TestArr(t *testing.T) {
	var a = make(map[string]HandlersChain)
	a["sdf"] = HandlersChain{}
	log.Println(a["sdf"] == nil)
}
