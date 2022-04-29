---
sidebar_label: Answer
sidebar_position: 1
---
<p align="center">
  <img width="150" height="174" src="https://raw.githubusercontent.com/fathiraz/kbs/main/ntf/logo.png"/>
</p>

# Notification Service
Notification service is a service for sending notifier to user, including :
- sms
- email (TODO)
- push notification (TODO)

## What you'll need
- [Golang](https://go.dev/dl//) version 1.17 or above
- Basic knowledge about go mod https://github.com/golang/go/wiki/Modules

## Prerequisites
Download all required library using
```bash
make download
```

Or you can use
```bash
go mod download
```

## Default configuration
You can change default configuration on file .env
```dotenv
# application config
APP_NAME=notification-service
APP_PORT=3000
APP_DEBUG=true

# authentication config
BASIC_USERNAME=kita
BASIC_PASSWORD=bisa

# sms vendor config
SMS_CONFIG_TYPE=toml
SMS_CONFIG_PATH=config/sms_vendor
```

- `APP_NAME` is configuration to set application name
- `APP_PORT` is configuration to set application api port, accept number
- `APP_DEBUG` is configuration to set application debug, accept only **true** / **false**
- `BASIC_USERNAME` is configuration to set basic username on request header
- `BASIC_PASSWORD` is configuration to set basic password on request header
- `SMS_CONFIG_TYPE` is configuration to set sms config type to use on the application, accept only **toml** / **yaml**, you can check example of configuration each file in [here](https://github.com/fathiraz/kbs/tree/main/ntf/config/sms_vendor).
- `SMS_CONFIG_PATH` is configuration to set sms config path to use on the application

## Build the app
You can build the app using one of these method, choose whats you prefer
- **direct running**
	```bash
	make run
	```
    Or you can use
	```bash
	go run main.go
	```
- **docker running**
	```bash
	make run-docker
	```
	Or you can use
	```bash
	docker-compose up
	```

## Running the test
The app is testable and use colorized output test by [gotest](https://github.com/rakyll/gotest). This test should be run with file `.env`, `config/sms_vendor/sms_vendor.yaml`, & `config/sms_vendor/sms_vendor.toml` has been set.

### Unit test
Basic test with no coverage.
```bash
make test
```

### Unit test + coverage
Test with coverage that will produced output to file `coverage-data.txt`.
```bash
make cover
```

## Running the app
You can build the app first on [this step](https://fathiraz.github.io/kbs/docs/problem-5/answer#build-the-app). Import postman environment and collection [right here](https://github.com/fathiraz/kbs/tree/main/ntf/doc).
- Postman collection `doc/Kitabisa.postman_collection.json`
- Postman environment `doc/Kitabisa.	postman_environment.json`

After running the app here is what you can do

### Send Sms
- method `POST`
- url `{{url}}/v1/sms/send`

request
```bash
curl --location --request POST '{{url}}/v1/sms/send' \
--header 'Authorization: Basic {{basic_authentication}}' \
--header 'Content-Type: application/json' \
--data-raw '{
	"number" : "628987654321",
	"message" : "test"
}'
```

success response
```bash
{
    "success": true,
    "code": 200,
    "message": "message send successfully"
}
```

failed response
```bash
{
    "success": false,
    "code": 400,
    "message": "code=400, message=Number wajib diisi",
    "error": {
        "message": "Number wajib diisi"
    }
}
```

### Toggle Vendor
- method `POST`
- url `{{url}}/v1/sms/toggle`

request
```bash
curl --location --request POST '{{url}}/v1/sms/toggle' \
--header 'Authorization: Basic {{basic_authentication}}' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name" : "sms_vendor_kita",
	"toggle" : true
}'
```

response
```bash
{
    "success": true,
    "code": 200,
    "message": "sms vendor toggled successfully",
    "data": {
        "name": "sms_vendor_kita",
        "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",
        "endpoint": "/vendor_sms_kita",
        "token": "Basic abc123def456",
        "is_default": false,
        "is_active": true,
        "timeout": "30s"
    }
}
```

### Get All Vendor
- method `GET`
- url `{{url}}/v1/sms`

request
```bash
curl --location --request GET '{{url}}/v1/sms' \
--header 'Authorization: Basic {{basic_authentication}}'
```

response
```bash
{
    "success": true,
    "code": 200,
    "message": "sms vendor get successfully",
    "data": [
        {
            "name": "sms_vendor_kita",
            "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",
            "endpoint": "/vendor_sms_kita",
            "token": "Basic abc123def456",
            "is_default": false,
            "is_active": true,
            "timeout": "30s"
        },
        {
            "name": "sms_vendor_bisa",
            "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",
            "endpoint": "/vendor_sms_bisa",
            "token": "Token abc123def456",
            "is_default": false,
            "is_active": false,
            "timeout": "30s"
        },
        {
            "name": "sms_vendor_kitabisa",
            "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",
            "endpoint": "/vendor_sms_kitabisa",
            "token": "Auth abc123def456",
            "is_default": true,
            "is_active": true,
            "timeout": "30s"
        }
    ]
}
```

## Source code
You can check the source code right here **https://github.com/fathiraz/kbs/tree/main/ntf**
