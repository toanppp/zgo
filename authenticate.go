package zgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	ErrorResp
}

func (z *zgo) RefreshAccessToken(ctx context.Context, refreshToken string) (GetAccessTokenResp, error) {
	payload := url.Values{}
	payload.Set("app_id", z.appID)
	payload.Set("grant_type", "refresh_token")
	payload.Set("refresh_token", refreshToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, getAccessTokenURL, bytes.NewBufferString(payload.Encode()))
	if err != nil {
		return GetAccessTokenResp{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("secret_key", z.appSecret)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return GetAccessTokenResp{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return GetAccessTokenResp{}, errors.New("request failed")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetAccessTokenResp{}, err
	}

	var data GetAccessTokenResp
	if err := json.Unmarshal(responseBody, &data); err != nil {
		return GetAccessTokenResp{}, err
	}

	if data.Error != 0 {
		return data, fmt.Errorf("request failed: %+v", data)
	}

	return data, nil
}
