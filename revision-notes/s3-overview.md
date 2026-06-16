# s3 buckets

---

### s3 buckets are infrastructure and they hold s3 objects

---

Important Concepts

- ### S3 Bucket Naming Rules

- ### S3 Bucket Restrictions and and Limitations

- ### S3 Bucket Types (flat vs general purpose)

- ### S3 Bucket Folders (Only for general purpose buckets)

- ### S3 Bucket Versioning

- ### S3 Bucket Encryption

- ### Static Website Hosting

---

Even though s3 is a globally available service but we need to specify a reigon while creating a bucket

---

### S3 Bucket Naming Rules:

![S3 bucket overview](./images/s3-overview.png)
![](images/2026-06-11-14-48-01.png)




### S3 Bucket Restrictions and and Limitations:
![](images/2026-06-11-14-49-40.png)

### S3 Bucket Types:
![](images/2026-06-11-14-52-40.png)

### S3 Bucket Folder:
![](images/2026-06-11-14-53-54.png)

### S3 Bucket Folder Example:
- You can see that the folder here is not actually a folder but a prefix that is actually an empty bucket ending with a slash.
![](images/2026-06-11-14-56-04.png)
- Therefore, we can also say that it is a flat hierarchy.

### S3 Objects Overview:
![](images/2026-06-11-14-58-49.png)

- S3 Entity Tag (ETag)
![](images/2026-06-11-15-01-38.png)

- CFN can be used in various ways but terraform provides more control over the objects in a sense that even an object can be dealt as an infrastructure while CFN can't do that.

### S3 checksum:
- Checksum reference ![](images/2026-06-11-16-17-35.png)
- ETags and Checksums are very different because an ETag is used to track the changes in a file from the perspective of the developer programmatically and Checksum is used to ensure data integrity of the file being uploaded or downloaded from the aws bucket.
- The most accurate way to look at it is that a Checksum is a mathematical concept used for integrity, while an ETag (Entity Tag) is an HTTP protocol feature used for web caching and concurrency control, which uses checksums under the hood.

### S3 Object - prefixes
![](images/2026-06-12-15-47-10.png)

### S3 Objects - metadata
![](images/2026-06-15-11-42-42.png)

![](images/2026-06-15-11-44-24.png)

![](images/2026-06-15-11-45-19.png)

# WORM (Write Once Read Many)
![](images/2026-06-15-12-18-25.png)

### S3 Object Lock
![](images/2026-06-15-12-19-41.png)

![](images/2026-06-15-12-20-04.png)

# S3 Bucket URI
![](images/2026-06-15-12-22-31.png)

# AWS S3 CLI
![](images/2026-06-15-12-24-43.png)

# AWS S3 Request Styles
![](images/2026-06-15-12-31-31.png)

# S3 DualStack Endpoints
![](images/2026-06-15-12-33-29.png)

# S3 Storage Classes Overview
![](images/2026-06-15-12-37-57.png)

![](images/2026-06-15-12-40-21.png)

![](images/2026-06-15-12-41-57.png)

![](images/2026-06-15-13-02-14.png)

# One Zone lowest latency bucket type
![](images/2026-06-15-13-24-47.png)

- S3 Express One Zone is ideal for compute-heavy applications processing millions of small objects per second right alongside AWS compute resources.

# S3 One Zone IA
![](images/2026-06-15-13-31-50.png)

# Glacier Storage Classes and Glacier Vault
- It uses Vaults instead of buckets
![](images/2026-06-15-13-34-38.png)

# Glacier Instant Retrieval
![](images/2026-06-15-13-37-15.png)

# Glacier Flexible Retrieval
![](images/2026-06-15-13-44-30.png)

# Glacier Deep Archive
![](images/2026-06-15-13-46-05.png)

# Intelligent Tiering
![](images/2026-06-15-13-48-06.png)

# Storage Class Comparison
![](images/2026-06-15-13-54-18.png)

# S3 Security Overview
![](images/2026-06-15-13-58-49.png)

## S3 Block Public Access
![](images/2026-06-15-14-01-01.png)

## S3 ACL (Access Control Lists)
![](images/2026-06-15-14-02-05.png)

## S3 bucket policies
![](images/2026-06-15-16-10-41.png)

## IAM policies and bucket policies difference
![](images/2026-06-16-14-25-44.png)

## Access Grants
![](images/2026-06-16-14-26-56.png)

## IAM Access Analyzer for S3
![](images/2026-06-16-14-28-10.png)

## Internetwork Traffic Privacy
![](images/2026-06-16-14-30-42.png)

## CORS
![](images/2026-06-16-14-32-41.png)

![](images/2026-06-16-14-34-49.png)

## S3 Encryption
![](images/2026-06-16-16-04-45.png)

![](images/2026-06-16-16-06-13.png)

![](images/2026-06-16-16-08-02.png)

![](images/2026-06-16-16-09-06.png)

![](images/2026-06-16-16-10-02.png)

![](images/2026-06-16-16-10-24.png)

![](images/2026-06-16-16-10-56.png)

![](images/2026-06-16-16-31-35.png)

failed attempt:
![](images/2026-06-16-16-38-17.png)

correct one:
![](images/2026-06-16-16-48-16.png)

![](images/2026-06-16-16-49-04.png)

