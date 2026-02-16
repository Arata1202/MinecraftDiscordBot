resource "aws_security_group" "minecraft_bot_sg" {
  name        = "minecraft_bot_sg"
  description = "minecraft_bot_sg"

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "minecraft_bot_sg"
  }
}
