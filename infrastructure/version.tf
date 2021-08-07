terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version "3.53.0"
    }
  }
}

provider "aws" {
  region = var.
  access_key = "<your_aws_access_key>"
  secret_key = "<your_aws_secret_key>"
}