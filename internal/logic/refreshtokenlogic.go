package logic

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-common/jwtutil"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *core.RefreshTokenReq) (*core.RefreshTokenResp, error) {
	if in.GetToken() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "token is required")
	}

	// 刷新只校验签名，不要求旧 token 未过期（参考 refresh_token_logic 语义）。
	claims := &jwtutil.Claims{}
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	token, err := parser.ParseWithClaims(in.GetToken(), claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(l.svcCtx.Config.AccessSecret), nil
	})
	if err != nil || !token.Valid || claims.UserId == "" {
		return nil, errorx.New(errorx.InvalidArgument, "invalid token")
	}

	q := l.svcCtx.Query
	u, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(claims.UserId)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}
	if err = checkUserStatus(u); err != nil {
		return nil, err
	}

	newToken, expire, err := signUserToken(l.ctx, l.svcCtx, u, "core_user_refresh_token")
	if err != nil {
		return nil, err
	}

	return &core.RefreshTokenResp{Token: newToken, Expire: expire}, nil
}
