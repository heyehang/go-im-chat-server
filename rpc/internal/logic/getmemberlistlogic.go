package logic

import (
	"context"

	"github.com/heyehang/go-im-grpc/chat_server"
	"go-im-chat-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberListLogic {
	return &GetMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberListLogic) GetMemberList(in *chat_server.GetMemberListReq) (*chat_server.GetMemberListResp, error) {
	// todo: add your logic here and delete this line

	return &chat_server.GetMemberListResp{}, nil
}
