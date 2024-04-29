resource "yandex_vpc_network" "dumpling-network" {
  name = var.network_name
}

resource "yandex_vpc_subnet" "dumpling-subnet" {
  name = var.subnet_name
  v4_cidr_blocks = [var.v4_cidr_blocks]
  zone           = var.zone
  network_id     = yandex_vpc_network.dumpling-network.id
}
