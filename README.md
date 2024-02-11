# TRANHUYDUCSEVEN - GOLANG ROCKSDB UBUNTU

### This repo is the module for building golang with rocksdb

- It contains grocksdb from [linxGnu/grocksdb](https://github.com/linxGnu/grocksdb)
- And also images for building rocksdb with golang in ubuntu
- The images (`tranhuyducseven/rocksdb:8.10.0-ubuntu`) include:
  - Ubuntu  latest
  - Golang 1.21
  - Rocksdb 8.10.0

### How to use

1. Import `github.com/tranhuyducseven/rocksdb` if you want to use rocksdb in your golang project

2. Use `tranhuyducseven/rocksdb` image to build your golang project with rocksdb

#### Example
```Dockerfile
FROM tranhuyducseven/rocksdb:8.10.0-ubuntu as builder

WORKDIR /app

ARG APP_NAME=abcxyz

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# # Build the Go app
RUN CGO_CFLAGS="-I/usr/include" \
    CGO_LDFLAGS="-L/usr/lib -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" \
    go build -o $APP_NAME .

EXPOSE 12345

CMD ["./abcxyz"]

```
