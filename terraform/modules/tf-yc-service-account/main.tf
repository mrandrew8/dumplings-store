resource "yandex_iam_service_account" "service-account-k8s" {
  name        = "service-account-k8s"
  description = "Сервисный аккаунт с ролью k8s.clusters.agent на каталог, в котором создается кластер Managed Service for Kubernetes. От его имени будут создаваться ресурсы, необходимые кластеру Managed Service for Kubernetes.
Сервисный аккаунт с ролью container-registry.images.puller на каталог с реестром Docker-образов. От его имени узлы будут скачивать из реестра необходимые Docker-образы."
  folder_id   = "b1gi8bi93n7ltsrefcih"
}

resource "yandex_resourcemanager_folder_iam_member" "k8s.clusters.agent" {
  folder_id   = "b1gi8bi93n7ltsrefcih"
  role        = "k8s.clusters.agent"
  member      = "serviceAccount:${yandex_iam_service_account.service-account-k8s.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "container-registry.images.puller" {
  folder_id   = "b1gi8bi93n7ltsrefcih"
  role        = "container-registry.images.puller"
  member      = "serviceAccount:${yandex_iam_service_account.service-account-k8s.id}"
}

resource "yandex_iam_service_account_static_access_key" "sa-static-key" {
 service_account_id = "${yandex_iam_service_account.service-account-k8s.id}"
 description        = "сервисный ключ"
 pgp_key            = "keybase:keybaseusername"
 }