  variable "network_id" {
    type = string
}

variable "subnet_id" {
    type = string
}

variable "security_group_ids" {
    type = string
}

variable "service_account_id" {
    type = string
}

variable "node_service_account_id" {
    type = string
}

variable "key_id" {
    type = string
}

variable "name" {
    default = "k8s-zonal"
    type = string
}

variable "zone" {
    default = "ru-central1-a"
    type = string
}

variable "public_ip" {
    default = true
    type = bool
}

