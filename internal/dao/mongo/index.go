package mongo

import (
	"context"
	"go-im-chat-server/internal/dao/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createMsgTableIndex(ctx context.Context) {
	msgIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "chat_id", Value: 1},
			bson.E{Key: "seq", Value: 1},
			bson.E{Key: "type", Value: 1},
			bson.E{Key: "sender", Value: 1},
		},
		Options: nil,
	}
	msgIDIndexName := "id"
	msgIDIndexUniq := true
	msgIDIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "id", Value: 1},
		},
		Options: &options.IndexOptions{
			Name:   &msgIDIndexName,
			Unique: &msgIDIndexUniq,
		},
	}
	idx := make([]mongo.IndexModel, 0, 2)
	idx = append(idx, msgIDIndex, msgIndex)
	_, err := db.Collection(model.MsgTableScheme).Indexes().CreateMany(ctx, idx)
	if err != nil {
		panic(err)
		return
	}
}

func createChatTableIndex(ctx context.Context) {
	chatIDIndexUniq := true
	chatIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "creator", Value: 1},
			bson.E{Key: "ctime", Value: 1},
			bson.E{Key: "type", Value: 1},
		},
	}
	idIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "1", Value: 1},
		},
		Options: &options.IndexOptions{Unique: &chatIDIndexUniq},
	}
	indexs := make([]mongo.IndexModel, 0, 2)
	indexs = append(indexs, chatIndex, idIndex)
	_, err := db.Collection(model.ChatTableScheme).Indexes().CreateMany(ctx, indexs)
	if err != nil {
		panic(err)
		return
	}
}

func createChatMsgIndex(ctx context.Context) {
	chatMsgIndexUniq := true
	chatMsgUniqIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "chat_id", Value: 1},
			bson.E{Key: "msg_id", Value: 1},
			bson.E{Key: "seq", Value: 1},
		},
		Options: &options.IndexOptions{
			Unique: &chatMsgIndexUniq,
		},
	}
	idIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "1", Value: 1},
		},
		Options: &options.IndexOptions{Unique: &chatMsgIndexUniq},
	}
	indexs := make([]mongo.IndexModel, 0, 2)
	indexs = append(indexs, chatMsgUniqIndex, idIndex)
	_, err := db.Collection(model.ChatMsgTableScheme).Indexes().CreateOne(ctx, chatMsgUniqIndex)
	if err != nil {
		panic(err)
		return
	}
}

func createChatMemberIndex(ctx context.Context) {
	uniq := true
	idIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "id", Value: 1},
		},
		Options: &options.IndexOptions{Unique: &uniq},
	}
	chatIDUserIDIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "chat_id", Value: 1},
			bson.E{Key: "user_id", Value: 1},
		},
		Options: &options.IndexOptions{Unique: &uniq},
	}
	userIDchatIDIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "user_id", Value: 1},
			bson.E{Key: "chat_id", Value: 1},
		},
	}
	chatIDRoleUtimeIndex := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "chat_id", Value: 1},
			bson.E{Key: "role", Value: 1},
			bson.E{Key: "utime", Value: 1},
		},
	}
	indexs := make([]mongo.IndexModel, 0, 2)
	indexs = append(indexs, chatIDUserIDIndex, userIDchatIDIndex, chatIDRoleUtimeIndex, idIndex)
	_, err := db.Collection(model.ChatMemberTableScheme).Indexes().CreateMany(ctx, indexs)
	if err != nil {
		panic(err)
		return
	}
}
