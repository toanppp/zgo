package zgo

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

const (
	getAccessTokenURL  = "https://oauth.zaloapp.com/v4/oa/access_token"
	getConversationURL = "https://openapi.zalo.me/v2.0/oa/conversation"
)

type Zgo interface {
	GetAppID() string
	EventSignature(data string, timestamp string) string
	RefreshAccessToken(ctx context.Context, refreshToken string) (GetAccessTokenResp, error)
	GetConversation(ctx context.Context, accessToken string, reqData GetConversationReq) (GetConversationResp, error)
}

type zgo struct {
	appID      string
	appSecret  string
	oaSecret   string
	httpClient *http.Client
}

type Option func(z *zgo)

func OptionHTTPClient(client *http.Client) func(*zgo) {
	return func(z *zgo) {
		z.httpClient = client
	}
}

func New(appID, appSecret, oaSecret string, opts ...Option) Zgo {
	z := &zgo{
		appID:      appID,
		appSecret:  appSecret,
		oaSecret:   oaSecret,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(z)
	}

	return z
}

func (z *zgo) GetAppID() string {
	return z.appID
}

func (z *zgo) EventSignature(data string, timestamp string) string {
	data = z.appID + data + timestamp + z.oaSecret
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}
