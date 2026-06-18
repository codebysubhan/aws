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

## S3 Bucket Key
![](images/2026-06-17-13-08-09.png)

## S3 Client Side Encryption
This is very different from the SSE-C as in SSE-C the server is applying the decrytion using the KMS server. But in this S3 CLient side encrytion the server is not involved instead the encrytion and decryption happens at the same end on client side.
![](images/2026-06-17-13-11-11.png)

## S3 Data Consistency
![](images/2026-06-18-12-38-51.png)

## S3 Object Replication
![](images/2026-06-18-12-40-53.png)

## S3 Versioning
![](images/2026-06-18-12-42-12.png)

## S3 Lifecycle
![](images/2026-06-18-12-44-00.png)

## S3 Transfer Acceleration
![](images/2026-06-18-12-44-55.png)

![](images/2026-06-18-12-45-27.png)

## S3 Presigned URLs
![](images/2026-06-18-12-46-12.png)

![](images/2026-06-18-12-46-59.png)

## S3 Access Points
![](images/2026-06-18-12-47-54.png)

![](images/2026-06-18-12-48-24.png)

## Multi Region Access Points
![](images/2026-06-18-12-49-25.png)

## S3 Object Lambda Access Points
![](images/2026-06-18-12-51-44.png)

## Mountpoint for Amazon S3
![](images/2026-06-18-12-53-24.png)

![](images/2026-06-18-12-54-12.png)

## Archived Objects
![](images/2026-06-18-12-55-48.png)

## Requesters Pay
![](images/2026-06-18-12-56-55.png)

![](images/2026-06-18-12-58-01.png)

![](images/2026-06-18-12-59-18.png)

![](images/2026-06-18-13-00-13.png)

## AWS marketplace for S3
![](images/2026-06-18-13-02-15.png)

## S3 Batch Operations
![](images/2026-06-18-13-03-21.png)

## Amazon S3 Inventory
![](images/2026-06-18-13-04-36.png)

## Amazon S3 Select
![](images/2026-06-18-13-06-51.png)

## S3 Event Notifications
![](images/2026-06-18-13-08-24.png)

## S3 Storage Class Analysis
![](images/2026-06-18-13-09-49.png)

## S3 Storage Lens
![](images/2026-06-18-13-11-01.png)

## S3 Static Website Hosting
![](images/2026-06-18-13-12-54.png)

## S3 Multi Part Upload
![](images/2026-06-18-13-13-42.png)

![](images/2026-06-18-13-14-25.png)

![](images/2026-06-18-13-15-22.png)

## S3 Byte Range Fetching
![](images/2026-06-18-13-16-37.png)

![](images/2026-06-18-13-17-35.png)

## S3 Interoperability
![](images/2026-06-18-13-19-43.png)
