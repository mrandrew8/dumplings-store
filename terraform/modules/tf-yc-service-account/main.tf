resource "yandex_iam_service_account" "k8s-account" {
  name        = "k8s-account"
  description = "K8S service account"
}

resource "yandex_resourcemanager_folder_iam_member" "k8s-clusters-agent" {
  # Сервисному аккаунту назначается роль "k8s.clusters.agent".
  folder_id = "b1gi8bi93n7ltsrefcih"
  role      = "k8s.clusters.agent"
  member    = "serviceAccount:${yandex_iam_service_account.k8s-account.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "vpc-public-admin" {
  # Сервисному аккаунту назначается роль "vpc.publicAdmin".
  folder_id = "b1gi8bi93n7ltsrefcih"
  role      = "vpc.publicAdmin"
  member    = "serviceAccount:${yandex_iam_service_account.k8s-account.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "images-puller" {
  # Сервисному аккаунту назначается роль "container-registry.images.puller".
  folder_id = "b1gi8bi93n7ltsrefcih"
  role      = "container-registry.images.puller"
  member    = "serviceAccount:${yandex_iam_service_account.k8s-account.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "encrypterDecrypter" {
  # Сервисному аккаунту назначается роль "kms.keys.encrypterDecrypter".
  folder_id = "b1gi8bi93n7ltsrefcih"
  role      = "kms.keys.encrypterDecrypter"
  member    = "serviceAccount:${yandex_iam_service_account.k8s-account.id}"
}