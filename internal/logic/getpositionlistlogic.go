package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionListLogic) GetPositionList(in *core.PositionListReq) (*core.PositionListResp, error) {
	q := l.svcCtx.Query
	conds := []gen.Condition{}
	if in.GetName() != "" {
		conds = append(conds, q.SysPosition.Name.Like("%"+in.GetName()+"%"))
	}
	if in.GetCode() != "" {
		conds = append(conds, q.SysPosition.Code.Like("%"+in.GetCode()+"%"))
	}
	if in.GetRemark() != "" {
		conds = append(conds, q.SysPosition.Remark.Like("%"+in.GetRemark()+"%"))
	}

	do := q.SysPosition.WithContext(l.ctx)
	if len(conds) > 0 {
		do = do.Where(conds...)
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
	list, err := do.Order(q.SysPosition.Sort, q.SysPosition.ID).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.PositionInfo, 0, len(list))
	for _, p := range list {
		data = append(data, toPositionInfo(p))
	}
	return &core.PositionListResp{Total: uint64(total), Data: data}, nil
}
