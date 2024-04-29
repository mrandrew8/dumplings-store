variable "cluster_id" {
    type = string
} 

variable "subnet_id" {
    type = string
}

variable "security_group_ids" {
    type = string
}

variable "name" {
    default = "k8s-node-group"
    type = string
}

variable "platform_id" {
    default = "standard-v2"
    type = string
}

variable "memory" {
    default = 2
    type = number
}

variable "cores" {
    default = 2
    type = number
}

variable "size" {
    default = 64
    type = number
}

variable "container_runtime" {
    default = "containerd"
    type = string
}

variable "min" {
    default = 1
    type = number
}

variable "max" {
    default = 3
    type = number
}

variable "initial" {
    default = 1
    type = number
}

variable "zone" {
    default = "ru-central1-a"
    type = string
}
