# SSM

ssm:
	@aws ssm start-session --target ${EC2_INSTANCE_ID}

up:
	@docker compose up -d

down:
	@docker compose down

.PHONY: ssm up down
