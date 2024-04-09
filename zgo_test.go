package zgo_test

import (
	"testing"

	"github.com/toanppp/zgo"
)

func TestApp_EventSignature(t *testing.T) {
	type fields struct {
		appID    string
		oaSecret string
	}
	type args struct {
		data      string
		timestamp string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "remove_user_from_tag",
			fields: fields{
				appID:    "appID",
				oaSecret: "oaSecret",
			},
			args: args{
				data:      `{"oa_id":"oa_id","user_id_by_app":"user_id_by_app","event_name":"remove_user_from_tag","tag":{"user_ids":["user_id"],"name":"Hoàn thành"},"app_id":"app_id","timestamp":"1702095007436"}`,
				timestamp: "1702095007436",
			},
			want: "315f0bb0c84fa87c816e302a590f5d26e6112a950c81bf28fb6b795130a19d09",
		},
		{
			name: "add_user_to_tag",
			fields: fields{
				appID:    "appID",
				oaSecret: "oaSecret",
			},
			args: args{
				data:      `{"oa_id":"oa_id","user_id_by_app":"user_id_by_app","event_name":"add_user_to_tag","tag":{"user_ids":["user_id"],"name":"Hoàn thành"},"app_id":"app_id","timestamp":"1702095015762"}`,
				timestamp: "1702095015762",
			},
			want: "f0866f93f29edf589f83fb9edb3b57c47fc34eef8bfd71dd6a00daa517382007",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := zgo.New(tt.fields.appID, "", tt.fields.oaSecret)
			if got := a.EventSignature(tt.args.data, tt.args.timestamp); got != tt.want {
				t.Errorf("EventSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
