FROM golang:1.22-alpine AS build

WORKDIR /src

COPY . ./

RUN go mod download && \
    go build -o /out/claude-telegram-bot

FROM alpine

COPY --from=build /out/claude-telegram-bot /bin

CMD [ "/bin/claude-telegram-bot" ]
