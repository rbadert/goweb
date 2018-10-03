package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"os"
)

func main() {
	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	//web.Use(logger.New())
	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		//Columns: true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKey: "logger_message",
	})

	app.Use(customLogger)
	app.StaticWeb("/node_modules", "./web/node_modules")
	// load templates
	app.RegisterView(iris.HTML("./web", ".html").Reload(true))

	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.ViewData("message", "The message from go web.")
		//ctx.HTML("<b>Welcome!</b>")
		ctx.View("views/index.html")
	})

	// same as web.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx context.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://<server>/hello
	app.Get("/hello", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Hello iris web framework."})
	})
	// or catch all http errors:
	app.OnAnyErrorCode(customLogger, func(ctx context.Context) {
		// this should be added to the logs, at the end because of the `logger.Config#MessageContextKey`
		ctx.Values().Set("logger_message","a dynamic message passed to the logs")
		ctx.Writef("My Custom error page")
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	var port string = os.Getenv("portnumber");
	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}