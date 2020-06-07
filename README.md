## iris学习笔记整理

参考文档：https://www.studyiris.com/doc/

### 1、iris安装

官方要求golang版本为1.8+，所以只要golang版本不是很久之前安装的老版本，基本没有问题

go get -u github.com/kataras/iris

创建main.go(文件名自己定义)

```go
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
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Run(iris.Addr(":8088"))
}
```

```shell
go run main.go
```

打开浏览器访问：127.0.0.1:8088/hello

可以看到控制台输出 Welcome则证明安装没有问题

### 2、关于iris如何管理配置

#### 1）可以通过内部配置：

app.Run函数第二个参数可以指定配置，例如：

```go
        app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{
            DisableInterruptHandler:           false,
            DisablePathCorrection:             false,
            EnablePathEscape:                  false,
            FireMethodNotAllowed:              false,
            DisableBodyConsumptionOnUnmarshal: false,
            DisableAutoFireStatusCode:         false,
            TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
            Charset:                           "UTF-8",
        }))
```

#### 2) 通过toml文件或者yaml文件管理配置

创建conf.toml

```toml
    DisablePathCorrection = false
    EnablePathEscape = false
    FireMethodNotAllowed = true
    DisableBodyConsumptionOnUnmarshal = false
    TimeFormat = "Mon, 01 Jan 2006 15:04:05 GMT"
    Charset = "UTF-8"
    [Other]
        MyServerName = "iris"
```

或者conf.yaml

```yaml
    DisablePathCorrection: false
    EnablePathEscape: false
    FireMethodNotAllowed: true
    DisableBodyConsumptionOnUnmarshal: true
    TimeFormat: Mon, 01 Jan 2006 15:04:05 GMT
    Charset: UTF-8
```

代码内部可以通过iris.TOML完成从配置文件到变量的转换

```go
    package main
    import (
       "github.com/kataras/iris"
    )
    func main() {
        app := iris.New()
        app.Get("/", func(ctx iris.Context) {
            ctx.HTML("<b>Hello!</b>")
        })
        // [...]
        // 通过文件配置 我们可以更加方便的切换开发环境配置和生产环境.
        app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./configs/iris.toml")))
        app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./configs/iris.yaml")))

    }
```

***

#### Built'n配置器

```go
// err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
// 当配置此项 如果web服务器 出现异常 我们将返回nil.
// 参考`Configuration的IgnoreServerErrors方法
// 地址: https://github.com/kataras/iris/tree/master/_examples/http-listening/listen-addr/omit-server-errors
func WithoutServerError(errors ...error) Configurator

// 当主服务器打开时，是否显示启动信息 如下
//Now listening on: http://localhost:8080
// Application started. Press CTRL+C to shut down.

var WithoutStartupLog

//当按下ctrl+C 时 禁止关闭当前程序(不会中止程序的运行)
var WithoutInterruptHandler

//路径重新定义(默认关闭)比如当访问/user/info 当该路径不存在的时候自动访问/user对应的handler
var WithoutPathCorrection

//如果此字段设置为true，则将创建一个新缓冲区以从请求主体读取。
var WithoutBodyConsumptionOnUnmarshal

//如果为true则关闭http错误状态代码处理程序自动执行
var WithoutAutoFireStatusCode

//转义路径
var WithPathEscape

//开启优化
var WithOptimizations

//不允许重新指向方法
var WithFireMethodNotAllowed

//设置时间格式
func WithTimeFormat(timeformat string) Configurator

//设值程序字符集
func WithCharset(charset string) Configurator

//启用或添加新的或现有的请求标头名称
func WithRemoteAddrHeader(headerName string) Configurator

//取消现有的请求标头名称
func WithoutRemoteAddrHeader(headerName string) Configurator

//自定义配置 key=>value
func WithOtherValue(key string, val interface{}) Configurator
```

***

