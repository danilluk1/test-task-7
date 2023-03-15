FROM alpine as base

RUN apk add git curl wget upx

WORKDIR /app

COPY --from=golang:1.20.1-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o ./bot ./cmd/main.go &&  \
  upx -9 -k ./bot
FROM base as bot
COPY --from=base /app/bot /bin/bot
CMD ["/bin/bot"]