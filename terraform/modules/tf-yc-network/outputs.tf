output "dumpling-network-id" {
  value = yandex_vpc_network.dumpling-network.id
}

output "dumpling-subnet-v4-cidr-blocks" {
  value = yandex_vpc_subnet.dumpling-subnet.v4_cidr_blocks
}

output "dumpling-subnet-id" {
  value = yandex_vpc_subnet.dumpling-subnet.id
}
