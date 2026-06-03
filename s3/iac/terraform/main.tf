terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.47.0"
    }
  }
}
# as it is commneted, the environment variables will take the priority
# provider "aws" {
#   # Configuration options
# }