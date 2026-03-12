package main

import (
	"fmt"
	"io"

	"github.com/go-amwk/core"
	"github.com/go-amwk/web"
)

func main() {
	app := web.Default()

	app.Use(func(ctx core.Context) error {
		body, err := ctx.Body()
		if err != nil {
			ctx.Write([]byte(fmt.Sprintf("Error reading body: %v\n", err)))
			return nil
		}
		defer body.Close()

		data, err := io.ReadAll(body)
		if err != nil {
			ctx.Write([]byte(fmt.Sprintf("Error reading body: %v\n", err)))
			return nil
		}

		ctx.Write(data)

		return nil
	})

	app.Start()
}
