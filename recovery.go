package houyi

import "fmt"

func Recovery() HandlerFunc {
	return func(context *Context) (err error) {
		index := context.index
		defer func() {
			if err := recover(); err != nil {
				_, _ = fmt.Fprintf(DefaultErrorWriter, "[panic] [%v] --> %v : %v\n", context.uri, nameOfFunction(context.handlers[index]), err)
			}
		}()
		return context.Next()
	}
}
