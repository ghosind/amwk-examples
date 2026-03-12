package main

import (
	"fmt"

	"github.com/go-amwk/core"
	"github.com/go-amwk/web"
)

func main() {
	app := web.Default()

	app.Use(func(ctx core.Context) error {
		val := ctx.Header("X-Custom-Header")
		ctx.Write([]byte(fmt.Sprintf("Header Value: %s\n", val)))
		return nil
	})

	app.Start()
}
