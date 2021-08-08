
variable "name" {
  description = "the name of your service"
}

variable "environment" {
  description = "the environment name (dev, staging, prod)"
}
variable "region" {
  description = "AWS region"
}
variable "ecr_repository_url" {
  description = "docker repository url"
}
variable "ecs_cluster_name" {
  description = "ecs cluster name"
}
variable "ecs_service_name" {
  description = "ecs service name"
}

variable "ecr_repository_arn" {
  description = "ecr arn"
}
variable "github_token" {
  description = "github token for code pipeline"
}
variable "github_owner" {
  description = "github owner for code pipeline"
}
variable "github_repo" {
  description = "github repository for code pipeline"
}
variable "github_branch" {
  description = "github branch for code pipeline"
}
