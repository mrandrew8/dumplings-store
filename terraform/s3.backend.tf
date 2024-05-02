# Описание бэкенда хранения состояния
terraform {
backend "s3" {
  endpoints = { 
  s3 = "https://storage.yandexcloud.net/"
  }

  bucket     = "dumpling-s3-bucket"
  region     = "ru-central1"
  key        = "terraform.tfstate"

  skip_region_validation      = true
  skip_credentials_validation = true
  skip_requesting_account_id  = true # необходимая опция Terraform для версии 1.6.1 и старше.
  skip_s3_checksum            = true # необходимая опция при описании бэкенда для Terraform 
  }
}
