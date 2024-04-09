package zgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	SrcOA = iota
	SrcUser
)

type GetConversationResp struct {
	Data Conversation `json:"data"`
	ErrorResp
}

type Conversation []Message

func (c Conversation) Profile() Profile {
	if len(c) == 0 {
		return Profile{}
	}
	switch c[0].Src {
	case SrcOA:
		return Profile{
			UserID:      c[0].ToID,
			DisplayName: c[0].ToDisplayName,
		}
	case SrcUser:
		return Profile{
			UserID:      c[0].FromID,
			DisplayName: c[0].FromDisplayName,
		}
	}
	return Profile{}
}

type Profile struct {
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
}

type Message struct {
	Src             int    `json:"src" bson:"src"`
	Time            int64  `json:"time" bson:"time"`
	SentTime        string `json:"sent_time" bson:"sent_time"`
	FromID          string `json:"from_id" bson:"from_id"`
	FromDisplayName string `json:"from_display_name" bson:"from_display_name"`
	FromAvatar      string `json:"from_avatar" bson:"from_avatar"`
	ToID            string `json:"to_id" bson:"to_id"`
	ToDisplayName   string `json:"to_display_name" bson:"to_display_name"`
	ToAvatar        string `json:"to_avatar" bson:"to_avatar"`
	MessageID       string `json:"message_id" bson:"message_id"`
	Type            string `json:"type" bson:"type"`
	Message         string `json:"message" bson:"message"`
	Links           []Link `json:"links,omitempty" bson:"links"`
	Thumb           string `json:"thumb" bson:"thumb"`
	URL             string `json:"url" bson:"url"`
	Description     string `json:"description" bson:"description"`
}

type Link struct {
	Title       string `json:"title" bson:"title"`
	URL         string `json:"url" bson:"url"`
	Thumb       string `json:"thumb" bson:"thumb"`
	Description string `json:"description" bson:"description"`
}

type LinkDescription struct {
	Caption string `json:"caption"`
	Phone   string `json:"phone"`
}

func (l Link) String() string {
	var data LinkDescription
	if err := json.Unmarshal([]byte(l.Description), &data); err != nil {
		return l.URL
	}

	switch l.URL {
	case "https://zalo.me":
		return data.Caption
	case "www.zaloapp.com":
		return fmt.Sprintf("Danh thiếp: %s - %s", l.Title, data.Phone)
	default:
		return l.URL
	}
}

type GetConversationReq struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset"`
	Count  int   `json:"count"`
}

func (z *zgo) GetConversation(ctx context.Context, accessToken string, reqData GetConversationReq) (GetConversationResp, error) {
	u, err := url.Parse(getConversationURL)
	if err != nil {
		return GetConversationResp{}, err
	}

	bytes, err := json.Marshal(reqData)
	if err != nil {
		return GetConversationResp{}, err
	}
	params := url.Values{}
	params.Set("data", string(bytes))
	u.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return GetConversationResp{}, err
	}

	req.Header.Set("access_token", accessToken)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return GetConversationResp{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return GetConversationResp{}, errors.New("request failed")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetConversationResp{}, err
	}

	var data GetConversationResp
	if err := json.Unmarshal(responseBody, &data); err != nil {
		return GetConversationResp{}, err
	}

	if data.Error != 0 {
		return data, fmt.Errorf("request failed: %+v", data)
	}

	return data, nil
}
