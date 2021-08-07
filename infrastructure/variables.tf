variable "aws_access_key" {
  type        = string
  description = "AWS Access Key"
}

variable "aws_secret_key" {
  type        = string
  description = "AWS Secret Key"
}

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