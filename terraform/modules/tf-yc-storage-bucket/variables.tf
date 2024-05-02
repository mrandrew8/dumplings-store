variable "service_account_name" {
  default = "s3-service-account"
  type = string
}

variable "folder_id" {
  default = "b1gi8bi93n7ltsrefcih"
  type = string
  description = "folder id"
}

variable "s3_bucket_name" {
  default = "dumpling-s3-bucket"
  type = string
}
