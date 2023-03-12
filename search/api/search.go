package main

import (
	"flag"
	"fmt"

	"go-zero-demo/search/api/internal/config"
	"go-zero-demo/search/api/internal/handler"
	"go-zero-demo/search/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/common/middleware"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//注册全局中间件
	server.Use(middleware.NewGlobalAuthMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
