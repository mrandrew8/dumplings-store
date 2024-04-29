output "static-public-ip-address" {
  value = yandex_vpc_address.static-public-ip.external_ipv4_address
}
