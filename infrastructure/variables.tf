variable "region" {
  type        = string
  description = "AWS Region"
}

variable "app" {
  type        = string
  description = "Application Name"
}

variable "environment" {
  type        = string
  description = "Application Environment"
}

variable "public_subnets" {
  description = "List of public subnets"
}

variable "private_subnets" {
  description = "List of private subnets"
}

variable "availability_zones" {
  description = "List of availability zones"
}

variable "github_owner" {
  description = "GitHub repository owner"
  default     = "stojce"
}
variable "github_token" {
  description = "GitHub repository owner"
  default     = "stojce"
}

variable "github_repo" {
  description = "GitHub repository name"
  default     = "static-web-example"
}

variable "task_memory" {
  description = "Task definition memory"
  default     = 256
}

variable "task_cpu" {
  description = "Task definition CPU"
  default     = 128
}

variable "port" {
  description = "Task definition memory"
  default     = 8888
}


