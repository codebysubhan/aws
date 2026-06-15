# Create a Bucket
```sh
aws s3 mb s3://class-fun-sa-12312
```

# Create a file

```sh
echo "Hello mars" > hello.txt
```

# Copy the file to s3
```sh
aws s3 cp hello.txt s3://class-fun-sa-12312
```

### Storage Class can be set at object level not bucket level
# change storage class to standard IA (Infrequent Access)
```sh
aws s3 cp hello.txt s3://class-fun-sa-12312 --storage-class STANDARD_IA
```

