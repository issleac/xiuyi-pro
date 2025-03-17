package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	helV1 "xiuyiPro/api/helloworld/v1"
	"xiuyiPro/internal/conf"
	"xiuyiPro/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	httpNet "net/http"
)

var (
	tSvr *service.TurtleService
	iSvr *service.IdiomService
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, turtle *service.TurtleService, idiom *service.IdiomService, logger log.Logger) *http.Server {
	router := gin.Default()
	// 使用kratos中间件
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
	router.Use(CORS())

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	httpSrv := http.NewServer(opts...)
	helV1.RegisterGreeterHTTPServer(httpSrv, greeter)

	httpSrv.HandlePrefix("/", router)
	tSvr = turtle
	iSvr = idiom
	outerRouter(router)

	return httpSrv
}

func outerRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	t := router.Group("/x/turtle")
	{
		t.GET("/list", listTurtles)
		t.POST("/set/batch", setBatchTurtles)
	}
	i := router.Group("/x/idiom")
	{
		i.GET("/get", getIdiom)
		i.POST("/set/batch", setIdioms)
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许访问的域名，这里可以根据实际情况修改，例如改为具体的前端域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 去除多余的空格
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.JSON(httpNet.StatusOK, gin.H{"message": "Options请求成功"})
			return
		}

		c.Next()
	}
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}
