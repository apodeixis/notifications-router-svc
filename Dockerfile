FROM golang:1.19-alpine3.18 as buildbase

WORKDIR /go/src/github.com/apodeixis/notifications-router-svc

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/notifications-router-svc main.go

FROM alpine:3.18

COPY --from=buildbase /usr/local/bin/notifications-router-svc /usr/local/bin/notifications-router-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["notifications-router-svc"]
