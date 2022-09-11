package mongo

import (
	"context"
	"go-im-chat-server/internal/dao/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMembersByChatID(ctx context.Context, chatID string, limit, skip int64) (uids []*model.ChatMember, err error) {
	cur, err := db.Collection(model.ChatMemberTableScheme).Find(ctx, bson.M{"chat_id": chatID},
		&options.FindOptions{
			Limit: &limit,
			Skip:  &skip,
			Projection: bson.M{
				"user_id": 1,
			},
		},
	)
	if err != nil {
		return
	}
	err = cur.All(ctx, &uids)
	if err != nil {
		return
	}
	return
}

func InsertChatMember(ctx context.Context, members []interface{}) (err error) {
	if len(members) <= 0 {
		return
	}
	_, err = db.Collection(model.ChatMemberTableScheme).InsertMany(ctx, members)
	if err != nil {
		return
	}
	return
}

//type Msg struct {
//	ID      primitive.ObjectID `json:"_" bson:"_id"`
//	MsgId   string             `json:"id" bson:"id"`
//	Sender  int64              `json:"sender" bson:"sender"`
//	ChatID  string             `json:"chat_id" bson:"chat_id"`
//	Content string             `json:"content" bson:"content"`
//	Seq     uint64             `json:"seq" bson:"seq"`
//	Type    uint16             `json:"type" bson:"type"`
//}

func InsertMsg(ctx context.Context, msgList []interface{}) (err error) {
	if len(msgList) <= 0 {
		return
	}
	_, err = db.Collection(model.MsgTableScheme).InsertMany(ctx, msgList)
	return
}
