terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.50.0"
    }
  }
}

resource "aws_s3_bucket" "default" {
#   bucket = "my-tf-test-bucket"
# it will randomly generate the bucket name as it is not being supplied here and commented out

#   tags = {
#     Name        = "My bucket"
#     Environment = "Dev"
#   }
# no need for metatags either for now
}

# s3 object: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_object

resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.default.id
  key    = "myfile.txt" # name of the object that we want
  source = "myfile.txt" # the local source of the object that we want to upload
  etag = filemd5("myfile.txt")
}
