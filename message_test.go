package zgo_test

import (
	"testing"

	"github.com/toanppp/zgo"
)

func TestLink_String(t *testing.T) {
	type fields struct {
		Title       string
		URL         string
		Thumb       string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "link",
			fields: fields{
				URL: "https://google.com",
			},
			want: "https://google.com",
		},
		{
			name: "number",
			fields: fields{
				Title:       "Name",
				URL:         "https://zalo.me",
				Description: "{\"phone\":\"767263039\",\"caption\":\"767263039\",\"qrCodeUrl\":\"https:\\/\\/qr-talk.zdn.vn\\/9\\/77327221\\/b7b2f11d85566c083547.jpg\"}",
			},
			want: "767263039",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := zgo.Link{
				Title:       tt.fields.Title,
				URL:         tt.fields.URL,
				Thumb:       tt.fields.Thumb,
				Description: tt.fields.Description,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
