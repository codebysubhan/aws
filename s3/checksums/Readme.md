## Create a new S3 bucket


```md
aws s3 mb s3://checksums-example-2495
```

## Create a file that we will do a checksum on

```md
echo "hello mars" > myfile.txt
```

## Get the checksum of a file for md5
```md
md5sum myfile.txt
```
## 410f1ae1c7dc3301898bceb8d7ae7fe3  myfile.txt

## Upload our s3 file

```md
aws s3 cp myfile.txt s3://checksums-example-2495
aws s3api head-object --bucket checksums-example-2495 --key myfile.txt
```
- can use aws s3 put-object as it would return the etag and some other information
- get-object vs head-object:
    - get-object downloads the file and head-object gets only the meta data including etag
- the etag returned was same as the checksum we calculated above
```md
hp@pop-os:~/Desktop/aws/s3/checksums$ aws s3api head-object --bucket checksums-example-2495 --key myfile.txt
{
    "AcceptRanges": "bytes",
    "LastModified": "Thu, 11 Jun 2026 11:33:33 GMT",
    "ContentLength": 11,
    "ETag": "\"410f1ae1c7dc3301898bceb8d7ae7fe3\"",
    "ContentType": "text/plain",
    "ServerSideEncryption": "AES256",
    "Metadata": {}
}
hp@pop-os:~/Desktop/aws/s3/checksums$ 
```

## Uploading an object using the put-object from s3api
link: https://docs.aws.amazon.com/cli/latest/reference/s3api/put-object.html
```sh

cksum myfile.txt
hp@pop-os:~/Desktop/aws/s3/checksums$ cksum myfile.txt 
4057472123 11 myfile.txt
```
```sh
aws s3api put-object \
--bucket="checksums-example-2495" \
--key="myfilecrc32.txt" \
--body="myfile.txt" \
--checksum-algorithm="CRC32" \
--checksum-crc32="4057472123"
```
- This above checksum is not working for some reason. Let's try python's method of calculating the checksum using CRC32 as algorithm.
```python
import zlib
# 1. Open and read the whole file into memory
with open("myfile.txt", "rb") as file:
    file_bytes = file.read()
# 2. Calculate the CRC32 checksum
crc_value = zlib.crc32(file_bytes) & 0xFFFFFFFF
# 3. Format to standard 8-character hex
crc_hex = f"{crc_value:08X}"
print(crc_hex)
---
A095AB54
```
```sh
aws s3api put-object \
--bucket="checksums-example-2495" \
--key="myfilecrc32.txt" \
--body="myfile.txt" \
--checksum-algorithm="CRC32" \
--checksum-crc32="A095AB54"
```
- Still invalid :(
- let's try another algorithm SHA256

```sh
hp@pop-os:~/Desktop/aws/s3/checksums$ sha256sum myfile.txt 
53d58e94e61b1c2a641dc52b402729f76c3832978e37f7f31ad1286ae32a796e  myfile.txt
```

```sh
aws s3api put-object \
--bucket="checksums-example-2495" \
--key="myfilesha256.txt" \
--body="myfile.txt" \
--checksum-algorithm="SHA256" \
--checksum-sha256="53d58e94e61b1c2a641dc52b402729f76c3832978e37f7f31ad1286ae32a796e"
```

## What went wrong

- `--checksum-sha256` maps to the `x-amz-checksum-sha256` header. Per AWS, that value must be the Base64-encoded raw 256-bit digest, not the hex string sha256sum prints.

You sent:
```
53d58e94e61b1c2a641dc52b402729f76c3832978e37f7f31ad1286ae32a796e ← hex (invalid)
```
AWS expects something like:
```
U9WOlOYbHCpkHcUrQCcp93xsOCl443978P3GvRKjKnl= ← base64 (valid)
```

- Same root cause as your CRC32 attempts in the Readme: cksum decimal and Python hex are also the wrong representation. AWS wants base64 of the binary checksum bytes for all of these headers.


```sh
hp@pop-os:~/Desktop/aws/s3/checksums$ ol dgst -sha256 -binary myfile
penssl dgst -sha256 -binary myfile.txt | base64
U9WOlOYbHCpkHcUrQCcp92w4MpeON/fzGtEoauMqeW4=
```
```sh
aws s3api put-object \
--bucket="checksums-example-2495" \
--key="myfilesha256.txt" \
--body="myfile.txt" \
--checksum-algorithm="SHA256" \
--checksum-sha256="U9WOlOYbHCpkHcUrQCcp92w4MpeON/fzGtEoauMqeW4="
```
