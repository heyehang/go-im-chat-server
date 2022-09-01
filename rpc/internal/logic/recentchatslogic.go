package logic

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecentChatsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecentChatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecentChatsLogic {
	return &RecentChatsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecentChatsLogic) RecentChats(in *chat_server.RecentChatsReq) (*chat_server.RecentChatsResp, error) {
	// todo: add your logic here and delete this line

	return &chat_server.RecentChatsResp{}, nil
}
