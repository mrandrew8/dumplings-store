resource "yandex_vpc_network" "dumpling-network" {
  name = "dumpling-network"
}

resource "yandex_vpc_subnet" "dumpling-subnet" {
  name = "dumpling-subnet"
  v4_cidr_blocks = ["10.1.0.0/16"]
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.dumpling-network.id
}
