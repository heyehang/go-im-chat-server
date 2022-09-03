package logic

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatMsgListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatMsgListLogic {
	return &GetChatMsgListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChatMsgListLogic) GetChatMsgList(in *chat_server.GetChatMsgListReq) (*chat_server.GetChatMsgListResp, error) {
	// todo: add your logic here and delete this line

	return &chat_server.GetChatMsgListResp{}, nil
}
