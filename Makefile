DR := npx dotenvx run --

# SSM

ssm:
	@aws ssm start-session --target ${EC2_INSTANCE_ID}

# Dotenvx

encrypt:
	@npx dotenvx encrypt

decrypt:
	@npx dotenvx decrypt

# Docker

DC := docker compose

up-f:
	${DR} ${DC} up -d --force-recreate

stop:
	${DR} ${DC} stop

# Terraform

tf-init:
	@cd terraform && terraform init

tf-plan:
	@cd terraform && terraform plan

tf-apply:
	@cd terraform && terraform apply

tf-destroy:
	@cd terraform && terraform destroy

.PHONY: ssm encrypt decrypt up-f stop tf-init tf-plan tf-apply tf-destroy
