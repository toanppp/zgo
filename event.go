package zgo

import (
	"strconv"
	"time"
)

const (
	AddUserToTag      = "add_user_to_tag"
	RemoveUserFromTag = "remove_user_from_tag"

	TagFinished = "Hoàn thành"
)

type Event struct {
	OAID        string  `json:"oa_id"`
	UserIDByApp string  `json:"user_id_by_app"`
	EventName   string  `json:"event_name"`
	Tag         UserTag `json:"tag"`
	AppID       string  `json:"app_id"`
	Timestamp   string  `json:"timestamp"`
}

func (e Event) CreatedAt() time.Time {
	msec, err := strconv.ParseInt(e.Timestamp, 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.UnixMilli(msec)
}

type UserTag struct {
	UserIDs []string `json:"user_ids"`
	Name    string   `json:"name"`
}
