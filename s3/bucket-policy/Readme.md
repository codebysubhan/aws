## Create a bucket

```sh
aws s3 mb s3://bucket-policy-sa-12321
```

## Create bucket policy
https://docs.aws.amazon.com/cli/latest/reference/s3api/put-bucket-policy.html
```sh
aws s3api put-bucket-policy --bucket bucket-policy-sa-12321 --policy file://policy.json
```

## In another account access the bucket
```sh
touch hello.txt
aws s3 cp hello.txt s3://bucket-policy-sa-12321
aws s3 ls s3://bucket-policy-sa-12321
```

## Cleanup
```sh
aws s3 rm s3://bucket-policy-sa-12321/hello.txt
aws s3 rb s3://bucket-policy-sa-12321
```

