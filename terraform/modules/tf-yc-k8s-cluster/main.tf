resource "yandex_kubernetes_cluster" "k8s-zonal" {
  name = "k8s-zonal"
  network_id = var.network_id
  master {
    master_location {
      zone      = "ru-central1-a"
      subnet_id = var.subnet_id
    }
    security_group_ids = [var.security_group_ids]
  }
  service_account_id      = var.service_account_id
  node_service_account_id = var.node_service_account_id
  kms_provider {
    key_id = var.key_id
  }
}
