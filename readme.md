# awscli s3 in minio

- 建立 configure : `aws configure --profile localhost`
  - 填入 access & secret key
- 列出 minio 本地端 bucket : `aws --endpoint-url http://localhost:9000 s3 ls --profile localhost`
- 列出 bucket 中的檔案 : `aws --endpoint-url http://localhost:9000 s3 ls s3://test3 --profile localhost`
  - `--recursive` 
  - `--human-readable` 
  - `--summarize`
- 其他操作 : https://docs.minio.io/docs/how-to-use-aws-sdk-for-go-with-minio-server

