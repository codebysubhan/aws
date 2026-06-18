## Create a bucket

```sh
aws s3 mb s3://encrypt-fun-sa-12321
```

## Create a file

```sh
echo "Hello marsss" > hello.txt
```

```sh
hp@pop-os:~/Desktop/aws/s3/encryption-client$ go run main.go 
2026/06/18 12:32:47 Reading local file contents from hello.txt...

--- PUT via Encryption Client ---
Status: Success (hello.txt contents securely uploaded!)

--- GET WITH KEY via Encryption Client ---
Result: Hello marsss

--- GET WITHOUT KEY via Standard S3 Client ---
Resulting Raw Cipher Text: [204 5 110 180 252 106 251 175 29 116 217 33 123 41 96 222 228 33 238 197 137 79 198 43 81 22 71 85 182 207 22 5 239 156 229 63 5 176 46 0 57]
```

## Download the file
```sh
aws s3 cp s3://encrypt-fun-sa-12321/secret downloaded_secret.bin
```

## Decrypt it locally
```sh
python3 -c '
import sys
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

key = b"a-very-secret-32-byte-long-key!!"
with open("downloaded_secret.bin", "rb") as f:
    data = f.read()

# Go prepended a 12-byte nonce at the front
nonce = data[:12]
ciphertext = data[12:]

aesgcm = AESGCM(key)
plaintext = aesgcm.decrypt(nonce, ciphertext, None)
print("Decrypted Content:", plaintext.decode("utf-8"))
'
```

## Clean Up
```sh
aws s3 rm s3://encrypt-fun-sa-12321/hello.txt
aws s3 rb s3://encrypt-fun-sa-12321
```