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

