# For development
FROM golang:1.21.7-alpine as dev

ENV TZ=Asia/Tokyo

ENV ROOT=/go/src/app
ENV ENV=development
ENV CGO_ENABLED=0

WORKDIR ${ROOT}

RUN apk update && apk add git

RUN go install github.com/cosmtrek/air@v1.49.0
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.18.0
RUN go install go.uber.org/mock/mockgen@v0.4.0

CMD ["air", "-c", ".air.toml"]

# For builder
FROM golang:1.21.7-alpine as builder

ENV TZ=Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY . ${ROOT}
ENV CGO_ENABLED=0

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -ldflags "-s -w" -o main ${ROOT}/main.go

# For production
FROM gcr.io/distroless/static-debian11:latest as prod

ENV TZ=Asia/Tokyo
ENV ROOT=/go/src/app

# goのビルドステージのイメージを参照し，debian環境にコピー
COPY --from=builder /go/bin/goose /goose
COPY db/ddl /migrations

COPY --from=builder ${ROOT}/main /main

COPY --from=busybox:uclibc /bin/sh /bin/sh
# COPY entrypoint.sh /entrypoint.sh

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
CMD ["server"]
