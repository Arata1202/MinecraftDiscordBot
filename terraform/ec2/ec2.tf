resource "aws_instance" "minecraft_bot_server" {
  ami                    = var.ami
  instance_type          = var.instance_type
  iam_instance_profile   = aws_iam_instance_profile.minecraft_bot_iam_instance_profile.name
  vpc_security_group_ids = [aws_security_group.minecraft_bot_sg.id]

  root_block_device {
    volume_type           = var.volume_type
    volume_size           = var.volume_size
    encrypted             = var.volume_encrypted
    kms_key_id            = var.volume_kms_key_id
    delete_on_termination = false
  }

  tags = {
    Name = "minecraft_bot_server"
  }
}
