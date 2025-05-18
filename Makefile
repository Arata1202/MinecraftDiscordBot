up:
	@docker compose up -d

down:
	@docker compose down

logs:
	@docker compose logs minecraft

.PHONY: up down logs
