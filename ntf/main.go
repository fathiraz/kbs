package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"ntf/config/environment"
	"ntf/config/sms_vendor"
	sms_delivery "ntf/internal/modules/sms/delivery"
	sms_usecase "ntf/internal/modules/sms/usecase"
	"ntf/pkg/shared/constant"
	middleware_utils "ntf/pkg/utils/middleware"
	sms_vendor_utils "ntf/pkg/utils/sms_vendor"
	validator_utils "ntf/pkg/utils/validator"
	"os"
	"os/signal"
	"time"
)

// main program
func main() {

	// parse configuration
	environment.ParseEnv()
	sms_vendor.ParseSmsVendorConfig()

	// set sms vendor holder
	smsVendorHolder, err := sms_vendor_utils.NewSmsVendorHolder(sms_vendor.GetSmsVendorConfig())
	if err != nil {
		log.Fatal(err)
	}

	// set usecase
	smsUsecase := sms_usecase.NewSmsUsecase(smsVendorHolder)

	// set delivery
	smsDelivery := sms_delivery.NewSmsDeliveryEcho(smsUsecase)

	// set new echo instance
	app := echo.New()

	// set validator
	app.Validator = &validator_utils.CustomValidator{Validator: validator.New()}

	// set echo middleware
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	app.Use(middleware_utils.MiddlewareBasicAuth(environment.GetEnv().BasicUsername, environment.GetEnv().BasicPassword))

	// set prefix routes
	appV1 := app.Group(constant.ApiV1)

	// set group
	smsDelivery.Mount(appV1)

	// start app
	port := fmt.Sprintf(":%s", environment.GetEnv().AppPort)
	go func() {
		if err = app.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the app")
		}
	}()

	log.Print(fmt.Sprintf("%s up and running on port %s!!!", environment.GetEnv().AppName, environment.GetEnv().AppPort))

	// wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
