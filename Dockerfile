FROM ubuntu:latest

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    gcc \
    g++ \
    wget \
    make \
    cmake \
    libbz2-dev \
    libsnappy-dev \
    zlib1g-dev \
    liblz4-dev \
    libzstd-dev \
    libgflags-dev \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*


# Install golang 

ARG GOLANG_VERSION=1.21.7

RUN wget https://dl.google.com/go/go${GOLANG_VERSION}.linux-amd64.tar.gz -O go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

ENV PATH=$PATH:/usr/local/go/bin


# Install RocksDB

ARG ROCKSDB_VERSION=8.10.0

ARG BUILD_PATH=/tmp/build

RUN mkdir -p $BUILD_PATH

RUN cd $BUILD_PATH && wget https://github.com/facebook/rocksdb/archive/v${ROCKSDB_VERSION}.tar.gz && tar xzf v${ROCKSDB_VERSION}.tar.gz && \
    cd rocksdb-${ROCKSDB_VERSION}/ && \
    make shared_lib &&\
    mkdir -p /usr/local/rocksdb/lib && \
    mkdir -p /usr/local/rocksdb/include && \
    cp librocksdb.so* /usr/local/rocksdb/lib && \
    cp librocksdb.so* /usr/lib/ && \
    cp -r include /usr/local/rocksdb/ && \
    cp -r include/* /usr/include/ && \
    cd $BUILD_PATH && rm -rf *
