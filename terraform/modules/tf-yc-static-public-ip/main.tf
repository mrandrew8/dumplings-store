resource "yandex_vpc_address" "static-public-ip" {
  name = "static-public-ip"
  deletion_protection = true
  folder_id = "b1gi8bi93n7ltsrefcih"
  external_ipv4_address {
    zone_id = "ru-central1-a"
  }
}