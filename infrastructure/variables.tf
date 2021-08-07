variable "aws_region" {
  type        = string
  description = "AWS Region"
  default     = "us-east-2"
}

variable "profile" {
  type        = string
  description = "AWS CLI Profile"
}

variable "service" {
  type        = string
  description = "Application Name"
}

variable "environment" {
  type        = string
  description = "Application Environment"
}

variable "cidr" {
  description = "The CIDR block for the VPC."
  default     = "10.0.0.0/16"
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

variable "task_cpu" {
  description = "Task Definition CPU"
  default     = 128
}

variable "task_memory" {
  description = "Task Definition Memory"
  default     = 256
}

variable "port" {
  description = "Service port"
}

