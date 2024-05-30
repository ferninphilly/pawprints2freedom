variable "api_gateway_name" {
  type = string
}

variable "environment" {
  type = string
}

variable "region" {
  type = string
}

variable "sqs_lambda_name" {
  type = string
}

variable "go_runtime" {
  type = string
}

variable "root_domain_name" {
  type    = string
  default = "pptf-web.co.uk"
}

variable "stack_name" {
  type = string
}

variable "account_id" {
  type = string
}
