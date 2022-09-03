package logic

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinToChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinToChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinToChatLogic {
	return &JoinToChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinToChatLogic) JoinToChat(in *chat_server.JoinToChatReq) (*chat_server.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &chat_server.EmptyResp{}, nil
}
