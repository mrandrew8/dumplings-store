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
} 

module "tf-yc-k8s-cluster" {
  source = "./modules/tf-yc-k8s-cluster"
} 



