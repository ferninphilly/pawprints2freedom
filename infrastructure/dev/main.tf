terraform {
  backend "s3" {
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "= 5.36.0"
    }
  }
}

provider "aws" {
  region              = local.region
  profile             = "pawprints2freedom"
  allowed_account_ids = [local.account_id]
}

locals {
  environment = var.environment
  region      = var.region
  account_id  = var.account_id
}

module "infrastructure_module" {
  region           = var.region
  environment      = var.environment
  account_id       = var.account_id
  sqs_lambda_name  = "sqs_consumer"
  api_gateway_name = "pptf_gateway"
  source           = "../infrastructure_module"
  go_runtime       = "provided.al2023"
  stack_name       = "pawprints-to-freedom"
}
