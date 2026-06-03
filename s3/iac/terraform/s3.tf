resource "aws_s3_bucket" "example" {
    # now it will randomize the bucket name
    # bucket="my-tf-test-bucket"
    tags={
        Name            = "My Bucket"
        Environment     = "Dev"
    }
}