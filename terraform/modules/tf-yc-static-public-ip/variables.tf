variable "name" {
  default = "static-public-ip"
  type = string
  description = "static public ip name"
}

variable "folder_id" {
  default = "b1gi8bi93n7ltsrefcih"
  type = string
  description = "folder id"
}

variable "deletion_protection" {
  default = true
  type = bool
  description = "deletion_protection"
}

variable "zone_id" {
  default = "ru-central1-a"
  type = bool
  description = "zone_id"
}
