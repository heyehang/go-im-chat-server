package work

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/heyehang/go-im-grpc/user_server"
	"github.com/heyehang/go-im-pkg/pulsarsdk"
	"github.com/panjf2000/ants/v2"
	"github.com/zeromicro/go-zero/zrpc"
	"go-im-chat-server/internal/config"
	"time"
)

type Work struct {
	pool    *ants.Pool
	c       config.Config
	userCli user_server.UserClient
}

func NewWork(c config.Config) *Work {
	w := new(Work)
	pool, err := ants.NewPool(c.WorkPoolSize, ants.WithNonblocking(true))
	if err != nil {
		panic(err)
		return nil
	}
	w.pool = pool
	w.userCli = user_server.NewUserClient(zrpc.MustNewClient(c.UserSrv, zrpc.WithTimeout(time.Second*5), zrpc.WithNonBlock()).Conn())
	return nil
}

func (w *Work) Start(ctx context.Context) {
	pulsarsdk.SubscribeMsg(ctx, w.c.Pulsar.Topic, func(message pulsar.Message, err error) {
		//data := message.Payload()
		// 找群成员
		// 组装消息
		// 写入topic
	})
}
