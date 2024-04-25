resource "yandex_kubernetes_node_group" "k8s-node-group" {
  cluster_id  = var.cluster_id
  name        = "k8s-node-group"


  instance_template {
    platform_id = "standard-v2"

    network_interface {
      nat                = true
      subnet_ids         = [var.subnet_id]
      security_group_ids = [var.security_group_ids]
    }

    resources {
      memory = 2
      cores  = 2
    }

    boot_disk {
      type = "network-hdd"
      size = 64
    }

    scheduling_policy {
      preemptible = false
    }

    container_runtime {
      type = "containerd"
    }
  }

  scale_policy {
    auto_scale {
      min     = 1
      max     = 3
      initial = 1
    }
  }

  allocation_policy {
    location {
      zone = "ru-central1-a"
    }
  }

  maintenance_policy {
    auto_upgrade = true
    auto_repair  = true

    maintenance_window {
      day        = "monday"
      start_time = "15:00"
      duration   = "3h"
    }
  }
}