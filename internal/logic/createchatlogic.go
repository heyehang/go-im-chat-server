package logic

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChatLogic {
	return &CreateChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateChatLogic) CreateChat(req *chat_server.CreateChatReq) (resp *chat_server.CreateChatResp, err error) {
	if req == nil {
		return
	}
	//todo 待实现
	return
}

func (l *CreateChatLogic) checkArg(req *chat_server.CreateChatReq) (err error) {

	return
}
