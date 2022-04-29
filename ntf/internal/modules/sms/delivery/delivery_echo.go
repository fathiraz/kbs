package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ntf/internal/modules/sms/model"
	sms_usecase "ntf/internal/modules/sms/usecase"
	"ntf/pkg/shared/result"
)

// DeliveryEcho delivery
type DeliveryEcho struct {
	smsUsecase sms_usecase.SmsUsecase
}

// NewSmsDeliveryEcho create new rest delivery
func NewSmsDeliveryEcho(smsUsecase sms_usecase.SmsUsecase) *DeliveryEcho {
	return &DeliveryEcho{
		smsUsecase: smsUsecase,
	}
}

// Mount delivery with root "/sms"
func (h *DeliveryEcho) Mount(root *echo.Group) {

	// set prefix
	sms := root.Group("/sms")

	// set method
	sms.GET("", h.GetAll)
	sms.POST("/send", h.Send)
	sms.POST("/toggle", h.Toggle)
}

// Send function delivery
func (h *DeliveryEcho) Send(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		message = "message send successfully"
	)

	// bind data
	var payload model.SmsSendRequest
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	// validate data
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	// set usecase
	err := h.smsUsecase.Send(ctx, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	return c.JSON(http.StatusOK, result.ResultSuccess(message, nil))
}

// Toggle function delivery
func (h *DeliveryEcho) Toggle(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		message = "sms vendor toggled successfully"
	)

	// bind data
	var payload model.SmsToggleRequest
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	// validate data
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	// set usecase
	data, err := h.smsUsecase.Toggle(ctx, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	return c.JSON(http.StatusOK, result.ResultSuccess(message, data))
}

// GetAll function delivery
func (h *DeliveryEcho) GetAll(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		message = "sms vendor get successfully"
	)

	// set usecase
	data, err := h.smsUsecase.GetAll(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.ResultError(err.Error(), err))
	}

	return c.JSON(http.StatusOK, result.ResultSuccess(message, data))
}
