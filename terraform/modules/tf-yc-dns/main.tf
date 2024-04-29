resource "yandex_dns_zone" "dumpling-zone1" {
  name        = var.name
  description = var.description

  labels = {
    label1 = var.label1
  }

  zone    = var.zone
  public  = var.public
}

