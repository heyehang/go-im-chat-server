package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/heyehang/go-im-grpc/chat_server"
	"github.com/heyehang/go-im-pkg/tlog"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go-im-chat-server/internal/config"
	"go-im-chat-server/internal/dao/mongo"
	"go-im-chat-server/internal/server"
	"go-im-chat-server/internal/svc"
	"go-im-chat-server/internal/work"
	"go-im-chat-server/pkg/pulsar"
	"go-im-chat-server/pkg/pyroscope"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	configFile = flag.String("f", "etc/chat.yaml", "the config file")
)

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	fileWriter := logx.Reset()
	writer, err := tlog.NewMultiWriter(fileWriter)
	logx.Must(err)
	logx.SetWriter(writer)
	cancelCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx := svc.NewServiceContext(c)
	pulsar.Init(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		chat_server.RegisterChatServer(grpcServer, server.NewChatServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	mongo.Init(cancelCtx, c)
	// 添加全局中间件
	//s.AddUnaryInterceptors(logic.Auth)
	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start()
	}()
	w := work.NewWork(c)
	wg.Add(1)
	go func() {
		defer wg.Done()
		w.Start(cancelCtx)
	}()
	pyroscope.Start(cancelCtx, wg, c.Name, c.PyroscopeAddr, nil, true)
	sig := make(chan os.Signal, 1)
	//syscall.SIGINT 线上记得加上这个信号 ctrl + c
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		s := <-sig
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			pyroscope.Closed()
			wg.Wait()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
