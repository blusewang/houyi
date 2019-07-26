package houyi

import (
	"errors"
	"github.com/blusewang/houyi"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	y := houyi.New()
	y.NewLayer("a").Use(func(context *houyi.Context) (err error) {
		log.Println("middle 1")
		return errors.New("test err in 1")
	}).Use(func(context *houyi.Context) (err error) {
		log.Println("middle 2")
		return
	}).NewLayer("b").Use(func(context *houyi.Context) (err error) {
		log.Println("middle 3")
		return
	}).Hit("",func(context *houyi.Context) (err error) {
		log.Println("hit")
		return
	})

	log.Println(y.Handle("a.b",[]byte("data")))
}

func TestArr(t *testing.T) {
	var a = make(map[string]houyi.HandlersChain)
	a["sdf"] = houyi.HandlersChain{}
	log.Println(a["sdf"] == nil)
}