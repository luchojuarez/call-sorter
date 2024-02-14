package callservice_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
)

func TestClient_FindByPhoneAndMonth(t *testing.T) {
	type fields struct {
		repo []callservice.Model
	}
	type args struct {
		phoneNumber string
		month       time.Month
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []callservice.Model
		wantErr bool
	}{
		{
			name: "sort",
			fields: fields{
				repo: []callservice.Model{
					{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116792944", Duration: 462, Date: time.Time{}.Add(5 * time.Second), CallType: callservice.CallTypeNational},
					{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116798953", Duration: 392, Date: time.Time{}.Add(1 * time.Second), CallType: callservice.CallTypeNational},
					{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+19116798053", Duration: 466, Date: time.Time{}.Add(10 * time.Second), CallType: callservice.CallTypeNational},
				},
			},
			args: args{
				phoneNumber: "+549XXXXXXXXXX",
				month:       time.January,
			},
			want: []callservice.Model{
				{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116798953", Duration: 392, Date: time.Time{}.Add(1 * time.Second), CallType: callservice.CallTypeNational},
				{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116792944", Duration: 462, Date: time.Time{}.Add(5 * time.Second), CallType: callservice.CallTypeNational},
				{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+19116798053", Duration: 466, Date: time.Time{}.Add(10 * time.Second), CallType: callservice.CallTypeNational},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := callservice.NewInMemoryClient(tt.fields.repo)
			got, err := cli.FindByPhoneAndMonthAndYear(context.Background(), tt.args.phoneNumber, tt.args.month, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindByPhoneAndMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.FindByPhoneAndMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
