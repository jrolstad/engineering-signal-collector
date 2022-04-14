variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-west-1"
}

variable "environment" {
  description = "Environment the infrasructure is for."

  type    = string
  default = "prd"
}

variable "container_image_id" {
  description = "Container image id to use for lambdas."

  type    = string
  default = "24"
}