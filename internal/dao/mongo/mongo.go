package mongo

import (
	"context"
	"github.com/heyehang/go-im-pkg/mongosdk"
	"go-im-chat-server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	cli *mongo.Client
	db  *mongo.Database
)

func Init(ctx context.Context, c config.Config) {
	mongosdk.InitMongo(ctx, c.Mongo.Addr, c.Mongo.User, c.Mongo.Pwd, c.Mongo.PoolSize)
	cli = mongosdk.GetClient()
	db = cli.Database(c.Mongo.Db)
	createIndex(ctx)
}

func GetClient() *mongo.Client {
	return mongosdk.GetClient()
}

func GetDB() *mongo.Database {
	return db
}

func createIndex(ctx context.Context) {
	createMsgTableIndex(ctx)
	createChatTableIndex(ctx)
	createChatMsgIndex(ctx)
	createChatMemberIndex(ctx)
}
