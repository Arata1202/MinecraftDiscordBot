variable "region" {
  default = "ap-northeast-1"
}

variable "ami" {
  # default = "ami-026c39f4021df9abe" # x86_64
  default = "ami-038e94aea55c0f480" # arm64
}

variable "instance_type" {
  default = "t4g.small"
}

variable "key_name" {
  default = "minecraft_key"
}

variable "volume_type" {
  default = "gp3"
}

variable "volume_size" {
  default = "8"
}
