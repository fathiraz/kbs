package usecase

import (
	"context"
	"ntf/config/sms_vendor"
	"ntf/internal/modules/sms/model"
	sms_vendor_utils "ntf/pkg/utils/sms_vendor"
)

type (

	// SmsUsecase abstraction
	SmsUsecase interface {
		Send(ctx context.Context, data *model.SmsSendRequest) (err error)
		Toggle(ctx context.Context, data *model.SmsToggleRequest) (result sms_vendor.SmsVendorConfig, err error)
		GetAll(ctx context.Context) (result []sms_vendor.SmsVendorConfig, err error)
	}

	smsUsecaseImpl struct {
		smsVendorHolder sms_vendor_utils.SmsVendorHolder
	}
)

// NewSmsUsecase usecase impl constructor
func NewSmsUsecase(smsVendorHolder sms_vendor_utils.SmsVendorHolder) SmsUsecase {
	return &smsUsecaseImpl{
		smsVendorHolder: smsVendorHolder,
	}
}

// Send function to send sms
func (s smsUsecaseImpl) Send(ctx context.Context, data *model.SmsSendRequest) (err error) {

	var (
		payload []byte
	)

	// get active vendor
	payload, err = s.setSendDataByActiveVendor(data)
	if err != nil {
		return
	}

	// send sms
	_, err = s.smsVendorHolder.ActiveVendor().Send(ctx, payload)
	if err != nil {
		return
	}

	return
}

// Toggle function to toggle sms
func (s smsUsecaseImpl) Toggle(ctx context.Context, data *model.SmsToggleRequest) (result sms_vendor.SmsVendorConfig, err error) {

	// get vendor
	vendor := s.smsVendorHolder.Get(data.Name)

	// toggle sms
	err = vendor.Toggle(ctx, data.Toggle)
	if err != nil {
		return
	}

	// set result
	result = s.smsVendorHolder.Get(data.Name).Data()

	return
}

// GetAll function to get all sms vendor
func (s smsUsecaseImpl) GetAll(ctx context.Context) (result []sms_vendor.SmsVendorConfig, err error) {

	// get all sms
	result = s.smsVendorHolder.GetAll()
	if err != nil {
		return
	}

	return
}

func (s smsUsecaseImpl) setSendDataByActiveVendor(data *model.SmsSendRequest) (payload []byte, err error) {

	var (
		activeVendor = s.smsVendorHolder.ActiveVendor()
	)

	switch activeVendor.Data().Name {
	case sms_vendor.SmsVendorNameKita:
		payloadData := data.ToSmsVendorKitaRequest()
		payload, err = payloadData.JSON()
	case sms_vendor.SmsVendorNameBisa:
		payloadData := data.ToSmsVendorBisaRequest()
		payload, err = payloadData.JSON()
	case sms_vendor.SmsVendorNameKitabisa:
		payloadData := data.ToSmsVendorKitabisaRequest()
		payload, err = payloadData.JSON()
	}

	if err != nil {
		return
	}
	return
}
