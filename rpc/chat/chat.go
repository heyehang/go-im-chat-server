// Code generated by goctl. DO NOT EDIT!
// Source: chat.proto

package chat

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateChatReq           = chat_server.CreateChatReq
	CreateChatResp          = chat_server.CreateChatResp
	EmptyResp               = chat_server.EmptyResp
	GetChatMsgListReq       = chat_server.GetChatMsgListReq
	GetChatMsgListResp      = chat_server.GetChatMsgListResp
	GetMemberListReq        = chat_server.GetMemberListReq
	GetMemberListResp       = chat_server.GetMemberListResp
	JoinToChatReq           = chat_server.JoinToChatReq
	Member                  = chat_server.Member
	Msg                     = chat_server.Msg
	RecentChatsReq          = chat_server.RecentChatsReq
	RecentChatsResp         = chat_server.RecentChatsResp
	RemoveMemberFromChatReq = chat_server.RemoveMemberFromChatReq

	Chat interface {
		CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*CreateChatResp, error)
		RecentChats(ctx context.Context, in *RecentChatsReq, opts ...grpc.CallOption) (*RecentChatsResp, error)
		JoinToChat(ctx context.Context, in *JoinToChatReq, opts ...grpc.CallOption) (*EmptyResp, error)
		RemoveMemberFromChat(ctx context.Context, in *RemoveMemberFromChatReq, opts ...grpc.CallOption) (*EmptyResp, error)
		GetChatMsgList(ctx context.Context, in *GetChatMsgListReq, opts ...grpc.CallOption) (*GetChatMsgListResp, error)
		GetMemberList(ctx context.Context, in *GetMemberListReq, opts ...grpc.CallOption) (*GetMemberListResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) CreateChat(ctx context.Context, in *CreateChatReq, opts ...grpc.CallOption) (*CreateChatResp, error) {
	client := chat_server.NewChatClient(m.cli.Conn())
	return client.CreateChat(ctx, in, opts...)
}

func (m *defaultChat) RecentChats(ctx context.Context, in *RecentChatsReq, opts ...grpc.CallOption) (*RecentChatsResp, error) {
	client := chat_server.NewChatClient(m.cli.Conn())
	return client.RecentChats(ctx, in, opts...)
}

func (m *defaultChat) JoinToChat(ctx context.Context, in *JoinToChatReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := chat_server.NewChatClient(m.cli.Conn())
	return client.JoinToChat(ctx, in, opts...)
}

func (m *defaultChat) RemoveMemberFromChat(ctx context.Context, in *RemoveMemberFromChatReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := chat_server.NewChatClient(m.cli.Conn())
	return client.RemoveMemberFromChat(ctx, in, opts...)
}

func (m *defaultChat) GetChatMsgList(ctx context.Context, in *GetChatMsgListReq, opts ...grpc.CallOption) (*GetChatMsgListResp, error) {
	client := chat_server.NewChatClient(m.cli.Conn())
	return client.GetChatMsgList(ctx, in, opts...)
}

func (m *defaultChat) GetMemberList(ctx context.Context, in *GetMemberListReq, opts ...grpc.CallOption) (*GetMemberListResp, error) {
	client := chat_server.NewChatClient(m.cli.Conn())
	return client.GetMemberList(ctx, in, opts...)
}
