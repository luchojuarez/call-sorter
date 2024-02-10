package invoice_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/luchojuarez/call-sorter/internal/domain/invoice"
	"github.com/luchojuarez/call-sorter/internal/domain/invoice/mocks"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/user"
)

//go:generate mockgen -destination=./mocks/call_repo.go -package=mocks github.com/luchojuarez/call-sorter/internal/domain/invoice CallRepository
//go:generate mockgen -destination=./mocks/invoice_repo.go -package=mocks github.com/luchojuarez/call-sorter/internal/domain/invoice InvoiceRepository
//go:generate mockgen -destination=./mocks/user_repo.go -package=mocks github.com/luchojuarez/call-sorter/internal/domain/invoice UserRepository

func TestProcessor_Generate(t *testing.T) {
	userPhoneNumber := "+549XXXXXXXXXX"
	defaultUser := user.Model{
		PhoneNumber: userPhoneNumber,
		Address:     "742 Evergreen Terrace",
		Name:        "Homer",
		Friends: []string{
			"+5491167930920",
			"+5491167920944",
			"+5491167980954",
			"+5491167980953",
			"+5491167980951",
			"+191167980953",
		},
		CountryCode: "AR",
	}
	type args struct {
		phoneNumber string
		month       time.Month
	}
	type callRepositoryMock struct {
		expectedPhoneNumber string
		expectedMonth       time.Month
		responseCalls       []callservice.Model
		err                 error
	}
	type userRepositoryMock struct {
		expectedPhoneNumber string
		err                 error
		user                user.Model
	}
	type serviceMocks struct {
		callRepositoryMock callRepositoryMock
		userRepositoryMock userRepositoryMock
	}
	tests := []struct {
		name         string
		args         args
		serviceMocks serviceMocks
		want         invoice.Model
		wantErr      bool
	}{
		{
			name: "10 firends call should return total amount zero",
			args: args{
				phoneNumber: userPhoneNumber,
				month:       time.November,
			},
			serviceMocks: serviceMocks{
				callRepositoryMock: callRepositoryMock{
					expectedPhoneNumber: userPhoneNumber,
					expectedMonth:       time.November,
					responseCalls: []callservice.Model{
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167920944", Duration: 462, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167980953", Duration: 392, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+191167980953", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeInternational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167920944", Duration: 462, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167980953", Duration: 392, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+191167980953", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeInternational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+191167980953", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeInternational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
					},
					err: nil,
				},
				userRepositoryMock: userRepositoryMock{
					expectedPhoneNumber: userPhoneNumber,
					user:                defaultUser,
				},
			},
			want: invoice.Model{
				Month: time.November,
				User:  invoice.FromUserRepo(defaultUser),
				Calls: []invoice.Call{
					{PhoneNumber: "+5491167920944", Duration: 462, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167980953", Duration: 392, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+191167980953", Duration: 466, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167920944", Duration: 462, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167980953", Duration: 392, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+191167980953", Duration: 466, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+191167980953", Duration: 466, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 0},
				},
				TotalInternationalSeconds: 1398,
				TotalNationalSeconds:      2203,
				TotalFriendsSeconds:       3601,
				Total:                     0,
			},
		},
		{
			name: "11 firends call should shuld pay for last call",
			args: args{
				phoneNumber: userPhoneNumber,
				month:       time.November,
			},
			serviceMocks: serviceMocks{
				callRepositoryMock: callRepositoryMock{
					expectedPhoneNumber: userPhoneNumber,
					expectedMonth:       time.November,
					responseCalls: []callservice.Model{
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167920944", Duration: 462, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167980953", Duration: 392, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+191167980953", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeInternational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167920944", Duration: 462, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167980953", Duration: 392, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+191167980953", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeInternational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+191167980953", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeInternational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+5491167930920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
					},
					err: nil,
				},
				userRepositoryMock: userRepositoryMock{
					expectedPhoneNumber: userPhoneNumber,
					user:                defaultUser,
				},
			},
			want: invoice.Model{
				Month: time.November,
				User:  invoice.FromUserRepo(defaultUser),
				Calls: []invoice.Call{
					{PhoneNumber: "+5491167920944", Duration: 462, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167980953", Duration: 392, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+191167980953", Duration: 466, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167920944", Duration: 462, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167980953", Duration: 392, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+191167980953", Duration: 466, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+191167980953", Duration: 466, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 0},
					{PhoneNumber: "+5491167930920", Duration: 165, Timestamp: time.Time{}, Amount: 2.5},
				},
				TotalInternationalSeconds: 1398,
				TotalNationalSeconds:      2368,
				TotalFriendsSeconds:       3766,
				Total:                     2.5, //just las call!
			},
		},

		{
			name: "10 national calls should cost 10 times constant",
			args: args{
				phoneNumber: userPhoneNumber,
				month:       time.November,
			},
			serviceMocks: serviceMocks{
				callRepositoryMock: callRepositoryMock{
					expectedPhoneNumber: userPhoneNumber,
					expectedMonth:       time.November,
					responseCalls: []callservice.Model{
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116792944", Duration: 462, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116798953", Duration: 392, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+19116798053", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116793920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116790944", Duration: 462, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116798953", Duration: 392, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+19116798053", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116793920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+19116798053", Duration: 466, Date: time.Time{}, CallType: callservice.CallTypeNational},
						{OriginNumber: "+549XXXXXXXXXX", RecipientNumber: "+549116793920", Duration: 165, Date: time.Time{}, CallType: callservice.CallTypeNational},
					},
					err: nil,
				},
				userRepositoryMock: userRepositoryMock{
					expectedPhoneNumber: userPhoneNumber,
					user:                defaultUser,
				},
			},
			want: invoice.Model{
				Month: time.November,
				User:  invoice.FromUserRepo(defaultUser),
				Calls: []invoice.Call{
					{PhoneNumber: "+549116792944", Duration: 462, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+549116798953", Duration: 392, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+19116798053", Duration: 466, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+549116793920", Duration: 165, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+549116790944", Duration: 462, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+549116798953", Duration: 392, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+19116798053", Duration: 466, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+549116793920", Duration: 165, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+19116798053", Duration: 466, Timestamp: time.Time{}, Amount: 2.5},
					{PhoneNumber: "+549116793920", Duration: 165, Timestamp: time.Time{}, Amount: 2.5},
				},
				TotalInternationalSeconds: 0,
				TotalNationalSeconds:      3601,
				TotalFriendsSeconds:       0,
				Total:                     25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			//CallRepository mocks
			callRepoMock := mocks.NewMockCallRepository(mockCtrl)
			callRepoMock.
				EXPECT().
				FindByPhoneAndMonthAndYear(
					gomock.Any(),
					tt.serviceMocks.callRepositoryMock.expectedPhoneNumber,
					tt.serviceMocks.callRepositoryMock.expectedMonth,
					2020,
				).
				Return(
					tt.serviceMocks.callRepositoryMock.responseCalls,
					tt.serviceMocks.callRepositoryMock.err,
				).
				Times(1)

			// UserRepositoryer repo.
			userRepoMock := mocks.NewMockUserRepository(mockCtrl)
			userRepoMock.
				EXPECT().
				GetByPhoneNumber(gomock.Any(), tt.serviceMocks.userRepositoryMock.expectedPhoneNumber).
				Return(
					tt.serviceMocks.userRepositoryMock.user,
					tt.serviceMocks.userRepositoryMock.err,
				).
				Times(1)

			p := invoice.NewProcessor(callRepoMock, userRepoMock)
			got, err := p.Generate(context.Background(), tt.args.phoneNumber, tt.args.month, 2020)
			if (err != nil) != tt.wantErr {
				t.Errorf("Processor.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Processor.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
