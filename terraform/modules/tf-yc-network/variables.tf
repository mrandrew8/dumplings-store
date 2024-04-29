variable "network_name" {
    default = "dumpling-network"
    type = string
}

variable "subnet_name" {
    default = "dumpling-subnet"
    type = string
}

variable "v4_cidr_blocks" {
    default = "10.1.0.0/16"
    type = string
}

variable "zone" {
    default = "ru-central1-a"
    type = string
}