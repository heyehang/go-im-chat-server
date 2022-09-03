package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MsgTableScheme        = "msg"
	ChatTableScheme       = "chat"
	ChatMsgTableScheme    = "chat_msg"
	ChatMemberTableScheme = "chat_member"
)

// 索引 _idx_chat_id_seq_type
// _uniq_id
type Msg struct {
	ID      primitive.ObjectID `json:"_" bson:"_id"`
	MsgId   string             `json:"id" bson:"id"`
	Sender  int64              `json:"sender" bson:"sender"`
	ChatID  string             `json:"chat_id" bson:"chat_id"`
	Content string             `json:"content" bson:"content"`
	Seq     uint64             `json:"seq" bson:"seq"`
	Type    uint16             `json:"type" bson:"type"`
}

// _idx_id_creator_ctime
type Chat struct {
	ID      primitive.ObjectID `bson:"_id"`
	ChatId  string             `json:"id" bson:"id"`
	Creator int64              `json:"creator" bson:"creator"`
	Ctime   int64              `json:"ctime" bson:"ctime"`
	Name    string             `json:"name" bson:"name"`
	Type    uint16             `json:"type" bson:"type"`
}

// uniq_chat_id_seq
type ChatMsg struct {
	ID        primitive.ObjectID `bson:"_id"`
	ChatMsgId string             `json:"id" bson:"id"`
	ChatId    string             `json:"chat_id" bson:"chat_id"`
	MsgID     string             `json:"msg_id" bson:"msg_id"`
	Seq       int64              `json:"seq" bson:"seq"`
}

// _uniq_chat_id_user_id
// _idx_user_id_chat_id
// _idx_chatid_role_utime
// _idx_ctime
// _idx_utime
type ChatMember struct {
	ID           primitive.ObjectID `bson:"_id"`
	ChatMemberID string             `json:"id" bson:"id"`
	ChatId       string             `json:"chat_id" bson:"chat_id"`
	UserID       int64              `json:"user_id" bson:"user_id"`
	CTime        int64              `json:"ctime" bson:"ctime"`
	UTime        int64              `json:"utime" bson:"utime"`
	Role         uint8              `json:"role" bson:"role"`
}

// im-work 需要用到的消息
type MsgReq struct {
	ServerKey    string   `json:"server_key"` // 服务唯一key ,目前暂未赋值，用于定向推送给指定im-work服务
	DeviceTokens []string `json:"device_tokens,omitempty"`
	// 消息体
	Msg []byte `protobuf:"json:"msg,omitempty"`
}
