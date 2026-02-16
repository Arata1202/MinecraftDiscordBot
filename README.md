<div align="right">

![GitHub License](https://img.shields.io/github/license/Arata1202/MinecraftDiscordBot)

</div>

## Getting Started

### Prepare Repository

```bash
# Local and VM

# Clone repository
git clone git@github.com:Arata1202/MinecraftDiscordBot.git
cd MinecraftDiscordBot

# Install dependencies
npm install
```

### Create Resources with Terraform

```bash
# Local

# Move to repository
cd MinecraftDiscordBot/terraform

# Prepare and edit variables file
cp variables.tf.example variables.tf
vi variables.tf

# Create resources
terraform init
terraform plan
terraform apply
```

### Connect AWS EC2 with SSM

```bash
# Local

# Move to repository
cd MinecraftDiscordBot

# Prepare and edit .envrc file
cp .envrc.example .envrc
vi .envrc

# Allow direnv to load variables
direnv allow .

# Connect to AWS EC2 via SSM
make ssm

# Switch to ubuntu user
sudo -iu ubuntu
```

```env
# Required
export EC2_INSTANCE_ID=<EC2_INSTANCE_ID>
```

### Set Up MinecraftDiscordBot Server

```bash
# VM

# Move to repository
cd MinecraftDiscordBot

# Set up Ubuntu
./ubuntu/setup.sh

# Remove existing .env file
rm -f .env

# Prepare and edit .env file
cp .env.example .env
vi .env

# Encrypt .env file
make encrypt

# Start server
sudo make up
```
