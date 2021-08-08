provider "aws" {
  profile = var.aws-profile
  region  = var.aws-region
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.52.0"
    }
  }
}

module "vpc" {
  source             = "./vpc"
  name               = var.name
  cidr               = var.cidr
  private_subnets    = var.private_subnets
  public_subnets     = var.public_subnets
  availability_zones = var.availability_zones
  environment        = var.environment
}

module "sg" {
  source         = "./sg"
  name           = var.name
  vpc_id         = module.vpc.id
  environment    = var.environment
  container_port = var.container_port
}

module "alb" {
  source              = "./alb"
  name                = var.name
  vpc_id              = module.vpc.id
  subnets             = module.vpc.public_subnets
  environment         = var.environment
  alb_security_groups = [module.sg.alb]
  health_check_path   = var.health_check_path
}

module "ecr" {
  source      = "./ecr"
  name        = var.name
  environment = var.environment
}

module "ecs" {
  source                      = "./ecs"
  name                        = var.name
  environment                 = var.environment
  region                      = var.aws-region
  subnets                     = module.vpc.private_subnets
  aws_alb_target_group_arn    = module.alb.aws_alb_target_group_arn
  ecs_service_security_groups = [module.sg.ecs-task-sg]
  container_port              = var.container_port
  container_cpu               = var.container_cpu
  container_memory            = var.container_memory
  service_desired_count       = var.service_desired_count
  aws_ecr_repository_url      = module.ecr.ecr_repository_url
}

module "codebuild" {
  source             = "./codebuild"
  name               = var.name
  environment        = var.environment
  region             = var.aws-region
  ecr_repository_url = module.ecr.ecr_repository_url
  ecr_repository_arn = module.ecr.ecr_repository_arn
  github_branch      = var.github_branch
  github_owner       = var.github_owner
  github_repo        = var.github_repo
  github_token       = var.github_token
  ecs_service_name   = module.ecs.ecs_service_name
  ecs_cluster_name   = module.ecs.ecs_cluster_name
}
