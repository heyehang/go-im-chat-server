package main

import (
	"context"
	"flag"
	"fmt"
	"go-im-chat-server/internal/work"
	"go-im-chat-server/pkg/pulsar"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/config"
	"go-im-chat-server/internal/server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	pulsar.Init(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		chat_server.RegisterChatServer(grpcServer, server.NewChatServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
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
		ctx1, cancel := context.WithCancel(context.Background())
		defer cancel()
		w.Start(ctx1)
	}()

	sig := make(chan os.Signal, 1)
	//syscall.SIGINT 线上记得加上这个信号 ctrl + c
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		s := <-sig
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			wg.Wait()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
