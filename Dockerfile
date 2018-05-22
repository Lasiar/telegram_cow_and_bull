FROM amd64/golang as builder
RUN mkdir -p /src/telegramGame
RUN GOPATH=/ go get github.com/go-telegram-bot-api/telegram-bot-api
ADD ./* /src/telegramGame
RUN cd /src/telegramGame
WORKDIR /src/telegramGame
RUN  GOPATH=/ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

FROM amd64/alpine
RUN mkdir /app
COPY --from=builder /src/telegramGame/main /app
ADD prod.config.json /app/conf.json
WORKDIR /app
CMD ["/app/main"]
