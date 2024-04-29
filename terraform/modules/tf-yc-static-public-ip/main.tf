resource "yandex_vpc_address" "static-public-ip" {
  name = var.name
  deletion_protection = var.deletion_protection
  folder_id = var.folder_id
  external_ipv4_address {
    zone_id = var.zone_id
  }
}