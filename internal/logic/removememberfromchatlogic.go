package logic

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMemberFromChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveMemberFromChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMemberFromChatLogic {
	return &RemoveMemberFromChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveMemberFromChatLogic) RemoveMemberFromChat(in *chat_server.RemoveMemberFromChatReq) (*chat_server.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &chat_server.EmptyResp{}, nil
}
