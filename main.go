package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	fmt.Println("vim-go")
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.HTML("<h1>{{.}}</h1>")
	})
	app.Run(iris.Addr(":8088"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
}
