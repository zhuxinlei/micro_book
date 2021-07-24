package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/zhuxinlei/micro_book/book/cmd/rpc/book"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/internal/config"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/internal/server"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/book.yaml", "the config file")

func main() {
	cc := logx.LogConf{
		ServiceName:         "book.rpc",
		Mode:                "file",
		TimeFormat:          "",
		Path:                "logs",
		Level:               "",
		Compress:            false,
		KeepDays:            0,
		StackCooldownMillis: 0,
	}
	logx.MustSetup(cc)
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewBookServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		book.RegisterBookServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
