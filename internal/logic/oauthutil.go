package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"golang.org/x/oauth2"
)

// providerEntry 缓存某个 provider 的 oauth2 配置与用户信息地址。
// 参考实现用包级 map 缓存但无并发保护，这里加读写锁避免并发竞态。
type providerEntry struct {
	config  oauth2.Config
	infoURL string
}

var (
	oauthProviderMu    sync.RWMutex
	oauthProviderCache = make(map[string]*providerEntry)
)

func dropOauthProviderCache(name string) {
	if name == "" {
		return
	}
	oauthProviderMu.Lock()
	delete(oauthProviderCache, name)
	oauthProviderMu.Unlock()
}

// loadOauthProvider 按 name 读取 provider，优先命中缓存。
func loadOauthProvider(ctx context.Context, svcCtx *svc.ServiceContext, name string) (*providerEntry, error) {
	oauthProviderMu.RLock()
	entry, ok := oauthProviderCache[name]
	oauthProviderMu.RUnlock()
	if ok {
		return entry, nil
	}

	q := svcCtx.Query
	p, err := q.SysOauthProvider.WithContext(ctx).Where(q.SysOauthProvider.Name.Eq(name)).First()
	if err != nil {
		return nil, err
	}

	entry = buildProviderEntry(p)
	oauthProviderMu.Lock()
	oauthProviderCache[name] = entry
	oauthProviderMu.Unlock()
	return entry, nil
}

func buildProviderEntry(p *model.SysOauthProvider) *providerEntry {
	return &providerEntry{
		config: oauth2.Config{
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   replaceOauthKeywords(p.AuthURL, p),
				TokenURL:  p.TokenURL,
				AuthStyle: oauth2.AuthStyle(p.AuthStyle),
			},
			RedirectURL: p.RedirectURL,
			Scopes:      strings.Fields(p.Scopes),
		},
		infoURL: p.InfoURL,
	}
}

// replaceOauthKeywords 与参考实现一致，支持认证地址中的占位符。
func replaceOauthKeywords(urlData string, p *model.SysOauthProvider) string {
	result := strings.ReplaceAll(urlData, "CLIENT_ID", p.ClientID)
	result = strings.ReplaceAll(result, "SECRET", p.ClientSecret)
	result = strings.ReplaceAll(result, "REDIRECT_URL", p.RedirectURL)
	result = strings.ReplaceAll(result, "SCOPE", p.Scopes)
	return result
}

// parseProviderFromState 从 state 中取 provider。
// 约定 state 形如 "<random>-<provider>"（参考实现用 Split(state, "-")[1]）；
// 这里取最后一段，兼容 provider 名本身不含 '-' 的常见用法。
func parseProviderFromState(state string) string {
	if state == "" {
		return ""
	}
	parts := strings.Split(state, "-")
	return strings.TrimSpace(parts[len(parts)-1])
}

type oauthUserInfo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	NickName    string `json:"nickName"`
	Nickname    string `json:"nickname"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Avatar      string `json:"avatar"`
	Mobile      string `json:"mobile"`
	Phone       string `json:"phone"`
	PhoneNumber string `json:"phoneNumber"`
}

func (u *oauthUserInfo) nickname() string {
	switch {
	case u.NickName != "":
		return u.NickName
	case u.Nickname != "":
		return u.Nickname
	case u.Name != "":
		return u.Name
	default:
		return ""
	}
}

func (u *oauthUserInfo) avatar() string {
	if u.Picture != "" {
		return u.Picture
	}
	return u.Avatar
}

func (u *oauthUserInfo) phone() string {
	switch {
	case u.Mobile != "":
		return u.Mobile
	case u.PhoneNumber != "":
		return u.PhoneNumber
	default:
		return u.Phone
	}
}

// fetchOauthUserInfo 用授权码换取 token 并拉取用户信息。
// auth_style: 1 -> 将 TOKEN 占位符替换为 access_token 放在 URL 上；
// 其它（含 2 与 0）-> Authorization: Bearer <access_token>。
func fetchOauthUserInfo(ctx context.Context, entry *providerEntry, code string) ([]byte, error) {
	if entry.infoURL == "" {
		return nil, fmt.Errorf("oauth provider info_url is empty")
	}

	reqCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	token, err := entry.config.Exchange(reqCtx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %w", err)
	}

	var req *http.Request
	if entry.config.Endpoint.AuthStyle == oauth2.AuthStyleInParams {
		u := strings.ReplaceAll(entry.infoURL, "TOKEN", token.AccessToken)
		req, err = http.NewRequestWithContext(reqCtx, http.MethodGet, u, nil)
	} else {
		req, err = http.NewRequestWithContext(reqCtx, http.MethodGet, entry.infoURL, nil)
		if err == nil {
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		}
	}
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get oauth user info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("failed to read oauth user info response: %w", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("oauth user info request failed: %s %s", resp.Status, string(body))
	}

	var probe map[string]any
	if err := json.Unmarshal(body, &probe); err != nil {
		return nil, fmt.Errorf("invalid oauth user info json: %w", err)
	}
	return body, nil
}
