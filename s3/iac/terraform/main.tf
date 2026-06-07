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
resource "aws_s3_bucket" "default" {
  
}

resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.default.id
  key = "myfile.txt"
  source = "myfile.txt"
  etag = filemd5("myfile.txt")
}