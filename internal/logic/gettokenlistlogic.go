package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *core.TokenListReq) (*core.TokenListResp, error) {
	q := l.svcCtx.Query

	// 与参考实现一致：有用户筛选条件时先定位用户，再按 uuid 查 token
	var uuids []string
	if in.GetUsername() != "" || in.GetNickname() != "" || in.GetEmail() != "" || in.GetUuid() != "" {
		conds := []gen.Condition{}
		if in.GetUuid() != "" {
			conds = append(conds, q.SysUser.ID.Eq(in.GetUuid()))
		}
		if in.GetUsername() != "" {
			conds = append(conds, q.SysUser.Username.Eq(in.GetUsername()))
		}
		if in.GetNickname() != "" {
			conds = append(conds, q.SysUser.Nickname.Like("%"+in.GetNickname()+"%"))
		}
		if in.GetEmail() != "" {
			conds = append(conds, q.SysUser.Email.Eq(in.GetEmail()))
		}
		users, err := q.SysUser.WithContext(l.ctx).Where(conds...).Find()
		if err != nil {
			return nil, err
		}
		for _, u := range users {
			uuids = append(uuids, u.ID)
		}
	}

	do := q.SysToken.WithContext(l.ctx)
	if len(uuids) > 0 {
		do = do.Where(q.SysToken.UUID.In(uuids...))
	}
	total, err := do.Count()
	if err != nil {
		return nil, err
	}

	page, pageSize := in.GetPage(), in.GetPageSize()
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	list, err := do.Order(q.SysToken.CreatedAt.Desc()).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.TokenInfo, 0, len(list))
	for _, t := range list {
		data = append(data, toTokenInfo(t))
	}
	return &core.TokenListResp{Total: uint64(total), Data: data}, nil
}
