FROM golang:1.23.2-alpine AS builder

COPY . /github.com/artyom-gaibovich/chat-service/source/

WORKDIR /github.com/artyom-gaibovich/chat-service/source/

RUN go mod download
RUN go build -o ./bin/chat_service cmd/note_grpc_server_v1/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/artyom-gaibovich/chat-service/source/bin/chat_service .

CMD ["./chat_service"]