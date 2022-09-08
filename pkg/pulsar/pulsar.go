package pulsar

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/heyehang/go-im-pkg/pulsarsdk"
	"go-im-chat-server/internal/config"

	"time"
)

func Init(conf config.Config) {
	err := pulsarsdk.Init(pulsar.ClientOptions{
		URL:                     conf.Pulsar.Url,
		ConnectionTimeout:       time.Second * time.Duration(conf.Pulsar.ConnectionTimeout),
		OperationTimeout:        time.Second * time.Duration(conf.Pulsar.OperationTimeout),
		MaxConnectionsPerBroker: conf.Pulsar.MaxConnectionsPerBroker,
	}, 10000)
	if err != nil {
		panic(err)
	}
}
