terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = ">= 1.26"
    }
  }
  required_version = ">= 1.4.0"
}