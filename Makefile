run:
	@docker compose -f config/compose.yml build
	@docker stack deploy -c config/compose.yml yasasvi-site
stop:
	@docker stack rm yasasvi-site
dev:
	@go run .
