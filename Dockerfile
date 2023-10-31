# syntax=docker/dockerfile:1

# Stage: Build
FROM golang:1.21.3-alpine3.18 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd cmd
COPY internal internal

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o /botchi-go ./cmd/bot

# Stage: Build release image with lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /botchi-go /botchi-go

ARG DISCORD_BOT_TOKEN
ARG DISCORD_BOT_LOG_GUILD_ID
ARG DISCORD_BOT_LOG_CHANNEL_ID
ARG CMD_BOT_DEBUG

ENV DISCORD_BOT_TOKEN ${DISCORD_BOT_TOKEN}
ENV DISCORD_BOT_LOG_GUILD_ID ${DISCORD_BOT_LOG_GUILD_ID}
ENV DISCORD_BOT_LOG_CHANNEL_ID ${DISCORD_BOT_LOG_CHANNEL_ID}
ENV CMD_BOT_DEBUG ${CMD_BOT_DEBUG}

ENTRYPOINT [ "/botchi-go" ]