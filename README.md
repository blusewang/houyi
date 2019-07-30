# houyi
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fblusewang%2Fhouyi.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fblusewang%2Fhouyi?ref=badge_shield)

router with middleware


# example

```go
	y := houyi.New()
	y.
		NewLayer("a").
		Use(func(context *houyi.Context) (err error) {
			log.Println("middle 1")
			return
		}).
		Use(func(context *houyi.Context) (err error) {
			log.Println("middle 2")
			return
		}).
		NewLayer("b").
		Use(func(context *houyi.Context) (err error) {
			log.Println("middle 3")
			return
		}).
		Hit("", func(context *houyi.Context) (err error) {
			log.Println("hit")
			return
		})

	log.Println(y.Handle("a.b", []byte("data")))
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fblusewang%2Fhouyi.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fblusewang%2Fhouyi?ref=badge_large)