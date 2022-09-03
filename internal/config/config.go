package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	WorkPoolSize int `json:",optional"`
	Pulsar       *Pulsar
	UserSrv      zrpc.RpcClientConf
	Mongo        *Mongo
}

type Pulsar struct {
	Url string `json:",optional"`
	// 单位秒
	ConnectionTimeout int `json:",optional"`
	// 单位秒
	OperationTimeout int `json:",optional"`
	// 最大连接数
	MaxConnectionsPerBroker int `json:",optional"`
	// 订阅topic
	Topic string `json:",optional"`
	// im-work 消费消息的topic
	WorkTopic string `json:",optional"`
}

type Mongo struct {
	Addr     string
	Db       string
	User     string `json:",optional"`
	Pwd      string `json:",optional"`
	PoolSize int    `json:",optional"`
}
