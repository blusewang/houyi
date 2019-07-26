# houyi
router with middleware


# example

```go
    y := houyi.New()
	y.NewLayer("a")
    .Use(func(context *houyi.Context) (err error) {
		log.Println("middle 1")
		return 
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
```