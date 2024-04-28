module "tf-yc-network" {
  source = "./modules/tf-yc-network"
} 

module "tf-yc-service-account" {
  source = "./modules/tf-yc-service-account"
} 

module "tf-yc-kms-symmetric-key" {
  source = "./modules/tf-yc-kms-symmetric-key"
} 

module "tf-yc-sgroups" {
  source = "./modules/tf-yc-sgroups"
  network_id  = module.tf-yc-network.dumpling-network-id
  v4_cidr_blocks    = module.tf-yc-network.dumpling-subnet-v4-cidr-blocks
} 

module "tf-yc-k8s-cluster" {
  source = "./modules/tf-yc-k8s-cluster"
  network_id = module.tf-yc-network.dumpling-network-id
  subnet_id = module.tf-yc-network.dumpling-subnet-id
  security_group_ids = module.tf-yc-sgroups.k8s-public-services-id
  service_account_id = module.tf-yc-service-account.k8s-account-id
  node_service_account_id = module.tf-yc-service-account.k8s-account-id
  key_id = module.tf-yc-kms-symmetric-key.kms-key
  depends_on = [module.tf-yc-service-account]
} 

module "tf-yc-node-group" {
  source = "./modules/tf-yc-node-group"
  cluster_id  = module.tf-yc-k8s-cluster.k8s-cluster-id
  subnet_id = module.tf-yc-network.dumpling-subnet-id
  security_group_ids = module.tf-yc-sgroups.k8s-public-services-id
} 

module "tf-yc-dns" {
  source = "./modules/tf-yc-dns"
  data = module.tf-yc-static-public-ip.static-public-ip-address
  depends_on = [module.tf-yc-static-public-ip]
} 

module "tf-yc-static-public-ip" {
  source = "./modules/tf-yc-static-public-ip"
} 









