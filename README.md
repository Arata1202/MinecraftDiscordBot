<div align="right">

![GitHub License](https://img.shields.io/github/license/Arata1202/MinecraftDiscordBot)

</div>

## Getting Started

### Create Resources on AWS EC2 with Terraform

```bash
# Clone repository
git clone git@github.com:Arata1202/MinecraftDiscordBot.git
cd MinecraftDiscordBot/.terraform

# Prepare and edit variables file
mv variables.example.tf variables.tf
vi variables.tf

# Create Resources
terraform init
terraform plan
terraform apply
```

### Setup on AWS EC2

```bash
# Clone repository
git clone https://github.com/Arata1202/MinecraftDiscordBot.git
cd MinecraftDiscordBot

# Prepare and edit .env file
mv .env.example .env
vi .env

# Set up EC2
./.linux/setup.sh

# Start the bot
sudo make up

# Stop the bot
sudo make down
```
