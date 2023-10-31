# Variables
## Container image tag
tag = test-231031-03

.PHONY: lint build

lint:
	@golangci-lint run

build:
	@docker build \
	-t botchi-go:$(tag) \
	--build-arg DISCORD_BOT_TOKEN=${DISCORD_BOT_TOKEN} \
	--build-arg DISCORD_BOT_LOG_GUILD_ID=${DISCORD_BOT_LOG_GUILD_ID} \
	--build-arg DISCORD_BOT_LOG_CHANNEL_ID=${DISCORD_BOT_LOG_CHANNEL_ID} \
	--build-arg CMD_BOT_DEBUG=${CMD_BOT_DEBUG} \
	.