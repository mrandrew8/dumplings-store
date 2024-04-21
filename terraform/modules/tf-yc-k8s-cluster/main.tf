resource "yandex_kubernetes_cluster" "k8s-zonal" {
  name = "k8s-zonal"
  network_id = module.tf-yc-network.dumpling-network-id
  master {
    master_location {
      zone      = "ru-central1-a"
      subnet_id = module.tf-yc-network.dumpling-subnet-id
    }
    security_group_ids = [module.tf-yc-sgroups.k8s-public-services-id]
  }
  service_account_id      = module.tf-yc-service-account.k8s-account-id
  node_service_account_id = module.tf-yc-service-account.k8s-account-id
  depends_on = [module.tf-yc-service-account]
  kms_provider {
    key_id = module.tf-yc-kms-symmetric-key.kms-key
  }
}
