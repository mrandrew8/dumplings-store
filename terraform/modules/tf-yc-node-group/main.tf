resource "yandex_kubernetes_node_group" "k8s-node-group" {
  cluster_id  = var.cluster_id
  name        = var.name


  instance_template {
    platform_id = var.platform_id

    network_interface {
      nat                = true
      subnet_ids         = [var.subnet_id]
      security_group_ids = [var.security_group_ids]
    }

    resources {
      memory = var.memory
      cores  = var.cores
    }

    boot_disk {
      type = "network-hdd"
      size = var.size
    }

    scheduling_policy {
      preemptible = false
    }

    container_runtime {
      type = var.container_runtime
    }
  }

  scale_policy {
    auto_scale {
      min     = var.min
      max     = var.max
      initial = var.initial
    }
  }

  allocation_policy {
    location {
      zone = var.zone
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