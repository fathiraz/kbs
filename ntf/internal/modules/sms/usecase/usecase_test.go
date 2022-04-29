package usecase

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"ntf/config/sms_vendor"
	"ntf/internal/modules/sms/model"
	mocks_sms_vendor_utils "ntf/mocks/pkg/utils/sms_vendor"
	"reflect"
	"testing"
)

// set testing usecase
type (
	TestingUsecase struct {
		Usecase           SmsUsecase
		smsVendorHolder   *mocks_sms_vendor_utils.SmsVendorHolder
		smsVendorKita     *mocks_sms_vendor_utils.SmsVendor
		smsVendorBisa     *mocks_sms_vendor_utils.SmsVendor
		smsVendorKitabisa *mocks_sms_vendor_utils.SmsVendor
	}
)

// initialize testing
func (s *TestingUsecase) Initialize(t *testing.T) (usecase *TestingUsecase) {

	s.smsVendorKita = &mocks_sms_vendor_utils.SmsVendor{}
	s.smsVendorBisa = &mocks_sms_vendor_utils.SmsVendor{}
	s.smsVendorKitabisa = &mocks_sms_vendor_utils.SmsVendor{}
	s.smsVendorHolder = &mocks_sms_vendor_utils.SmsVendorHolder{}

	// set usecase
	s.Usecase = NewSmsUsecase(s.smsVendorHolder)
	return
}

func Test_smsUsecaseImpl_Send(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *model.SmsSendRequest
	}

	tests := []struct {
		name       string
		expectFunc func(t *TestingUsecase)
		args       args
		wantErr    error
	}{
		{
			name: "Success Case" +
				" sms_vendor_kita",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("ActiveVendor").Return(t.smsVendorKita)
				t.smsVendorKita.On("Data").Return(sms_vendor.SmsVendorConfig{
					Name: sms_vendor.SmsVendorNameKita,
				})
				t.smsVendorKita.On("Send", mock.Anything, mock.Anything).Return(nil, nil)
			},
			args: args{
				data: &model.SmsSendRequest{},
			},
			wantErr: nil,
		},
		{
			name: "Success Case" +
				" sms_vendor_bisa",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("ActiveVendor").Return(t.smsVendorBisa)
				t.smsVendorBisa.On("Data").Return(sms_vendor.SmsVendorConfig{
					Name: sms_vendor.SmsVendorNameBisa,
				})
				t.smsVendorBisa.On("Send", mock.Anything, mock.Anything).Return(nil, nil)
			},
			args: args{
				data: &model.SmsSendRequest{},
			},
			wantErr: nil,
		},
		{
			name: "Success Case" +
				" sms_vendor_kitabisa",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("ActiveVendor").Return(t.smsVendorKitabisa)
				t.smsVendorKitabisa.On("Data").Return(sms_vendor.SmsVendorConfig{
					Name: sms_vendor.SmsVendorNameKitabisa,
				})
				t.smsVendorKitabisa.On("Send", mock.Anything, mock.Anything).Return(nil, nil)
			},
			args: args{
				data: &model.SmsSendRequest{},
			},
			wantErr: nil,
		},
		{
			name: "Failed Case" +
				" on.smsVendorHolder.ActiveVendor().Send",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("ActiveVendor").Return(t.smsVendorKitabisa)
				t.smsVendorKitabisa.On("Data").Return(sms_vendor.SmsVendorConfig{
					Name: sms_vendor.SmsVendorNameKitabisa,
				})
				t.smsVendorKitabisa.On("Send", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("test"))
			},
			args: args{
				data: &model.SmsSendRequest{},
			},
			wantErr: fmt.Errorf("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = context.Background()

			testing := &TestingUsecase{}
			testing.Initialize(t)

			if tt.expectFunc != nil {
				tt.expectFunc(testing)
			}

			if gotResult := testing.Usecase.Send(tt.args.ctx, tt.args.data); !reflect.DeepEqual(gotResult, tt.wantErr) {
				t.Errorf("Send() = %v, want %v", gotResult, tt.wantErr)
			}
		})
	}
}

func Test_smsUsecaseImpl_Toggle(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *model.SmsToggleRequest
	}
	tests := []struct {
		name       string
		expectFunc func(t *TestingUsecase)
		args       args
		wantResult sms_vendor.SmsVendorConfig
		wantErr    error
	}{
		{
			name: "Success Case",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("Get", mock.Anything).Return(t.smsVendorKita)
				t.smsVendorKita.On("Toggle", mock.Anything, mock.Anything).Return(nil)
				t.smsVendorKita.On("Data").Return(sms_vendor.SmsVendorConfig{
					Name: sms_vendor.SmsVendorNameKita,
				})
			},
			args: args{
				data: &model.SmsToggleRequest{},
			},
			wantResult: sms_vendor.SmsVendorConfig{
				Name: sms_vendor.SmsVendorNameKita,
			},
			wantErr: nil,
		},
		{
			name: "Failed Case" +
				"on vendor.Toggle",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("Get", mock.Anything).Return(t.smsVendorKita)
				t.smsVendorKita.On("Toggle", mock.Anything, mock.Anything).Return(fmt.Errorf("test"))
			},
			args: args{
				data: &model.SmsToggleRequest{},
			},
			wantErr: fmt.Errorf("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = context.Background()

			testing := &TestingUsecase{}
			testing.Initialize(t)

			if tt.expectFunc != nil {
				tt.expectFunc(testing)
			}

			gotResult, err := testing.Usecase.Toggle(tt.args.ctx, tt.args.data)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Toggle() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Toggle() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_smsUsecaseImpl_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		expectFunc func(t *TestingUsecase)
		args       args
		wantResult []sms_vendor.SmsVendorConfig
		wantErr    error
	}{
		{
			name: "Success Case",
			expectFunc: func(t *TestingUsecase) {
				t.smsVendorHolder.On("GetAll").Return([]sms_vendor.SmsVendorConfig{
					sms_vendor.SmsVendorConfig{},
				})
			},
			args: args{},
			wantResult: []sms_vendor.SmsVendorConfig{
				sms_vendor.SmsVendorConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = context.Background()

			testing := &TestingUsecase{}
			testing.Initialize(t)

			if tt.expectFunc != nil {
				tt.expectFunc(testing)
			}

			gotResult, err := testing.Usecase.GetAll(tt.args.ctx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetAll() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
