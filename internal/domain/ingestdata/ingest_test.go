package ingestdata

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/luchojuarez/call-sorter/internal/domain/ingestdata/mocks"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -destination=./mocks/call_repo.go -package=mocks github.com/luchojuarez/call-sorter/internal/domain/ingestdata CallRepository

var validCsvInput = `numero origen,numero destino,duracion,fecha
+5491167980950,+191167980952,462,2020-11-10T04:02:45Z`

func TestClient_ReadAll(t *testing.T) {
	type args struct {
		input string
	}
	type saveCallRepoMock struct {
		param callservice.Model
		err   error
	}
	tests := []struct {
		name              string
		args              args
		saveCallRepoMocks []saveCallRepoMock
		err               error
	}{
		{
			name: "success case should return nil error",
			saveCallRepoMocks: []saveCallRepoMock{
				{
					param: callservice.Model{
						OriginNumber:    "+5491167980950",
						RecipientNumber: "+191167980952",
						Duration:        462,
						Date:            time.Date(2020, 11, 10, 4, 2, 45, 0, time.UTC),
						CallType:        "NATIONAL",
					},
					err: nil,
				},
			},
			args: args{
				input: validCsvInput,
			},
			err: nil,
		},
		{
			name: "repository error should return error",
			saveCallRepoMocks: []saveCallRepoMock{
				saveCallRepoMock{
					param: callservice.Model{
						OriginNumber:    "+5491167980950",
						RecipientNumber: "+191167980952",
						Duration:        462,
						Date:            time.Date(2020, 11, 10, 4, 2, 45, 0, time.UTC),
						CallType:        callservice.CallTypeInternational,
					},
					err: errors.New("save error"),
				},
			},
			args: args{
				input: validCsvInput,
			},
			err: errors.New("error saving call '{OriginNumber:+5491167980950 RecipientNumber:+191167980952 Duration:462 Date:2020-11-10 04:02:45 +0000 UTC CallType:NATIONAL}' at: 1"),
		},
		{
			name: "(1)invalid csv input format should return error",
			args: args{
				input: "i'm not a csv\ni'm a monster!",
			},
			err: errors.New("invalid row: 1"),
		},
		{
			name: "(2)invalid csv input format should return error",
			args: args{
				input: "i'm not a csv i'm a monster!",
			},
			err: errors.New("invalid_CSV_input"),
		},
		{
			name: "invalid call duration format should return error",
			args: args{
				input: "numero origen,numero destino,duracion,fecha\n" +
					"+5491167980950,+191167980952,nan,2020-11-10T04:02:45Z",
			},
			err: errors.New("invalid call duration format 'nan' at: 1"),
		},
		{
			name: "invalid call date (same RFC but with nano) format should return error",
			args: args{
				input: "numero origen,numero destino,duracion,fecha\n" +
					"+5491167980950,+191167980952,150,2006-01-02T15:04:05.999999999Z07:00",
			},
			err: errors.New("invalid call date format '2006-01-02T15:04:05.999999999Z07:00' at: 1"),
		},
		{
			name: "multipl error should return all of this joint",
			args: args{
				input: "numero origen,numero destino,duracion,fecha\n" +
					"+5491167980950,+191167980952,nan,2020-11-10T04:02:45Z\n" +
					"+5491167980950,+191167980952,150,2006-01-02T15:04:05.999999999Z07:00",
			},
			err: errors.New("invalid call duration format 'nan' at: 1\ninvalid call date format '2006-01-02T15:04:05.999999999Z07:00' at: 2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			callStorageMock := mocks.NewMockCallRepository(mockCtrl)

			for _, callMock := range tt.saveCallRepoMocks {
				callStorageMock.
					EXPECT().
					Save(gomock.Any(), callMock.param).
					Return(callMock.err).
					Times(1)
			}

			reader := strings.NewReader(tt.args.input)

			c := NewClient(callStorageMock)
			err := c.ReadAll(context.Background(), reader)
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
