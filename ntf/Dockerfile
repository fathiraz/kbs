# alpine linux golang
FROM golang:1.17-alpine as builder

# set workdir
WORKDIR /app

# copy data
COPY . .

# update data
RUN apk update
RUN apk add --no-cache ca-certificates

# run go command
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -a -tags -ldflags="-w -s" .

FROM busybox

# set workdir busybox
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/ntf .
COPY --from=builder /app/config/sms_vendor/sms_vendor.toml ./config/sms_vendor/
COPY --from=builder /app/config/sms_vendor/sms_vendor.yaml ./config/sms_vendor/
COPY --from=builder /app/.env .

EXPOSE 3000

ENV TZ Asia/Jakarta

CMD [ "./ntf" ]
