resource "yandex_kubernetes_cluster" "k8s-zonal" {
  name = var.name
  network_id = var.network_id
  master {
    master_location {
      zone      = var.zone
      subnet_id = var.subnet_id
    }
    public_ip = var.public_ip
    security_group_ids = [var.security_group_ids]
  }
  service_account_id      = var.service_account_id
  node_service_account_id = var.node_service_account_id
  kms_provider {
    key_id = var.key_id
  }
}
