package houyi

import (
	"fmt"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) (err error) {
		start := time.Now()

		// Process request
		err = c.Next()

		// time cost path err
		now := time.Now()
		_, _ = fmt.Fprintf(DefaultErrorWriter, "[HouYi] %v [%v] %v %v\n", now.Format("2006-01-02 15:04:05"), now.Sub(start), c.uri, err)

		return
	}
}
