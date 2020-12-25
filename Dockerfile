FROM golang
WORKDIR /app

# System deps
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool vim

# Protobuf
ENV PB_REL="https://github.com/protocolbuffers/protobuf/releases"
RUN curl -LO $PB_REL/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
RUN unzip protoc-3.11.4-linux-x86_64.zip -d /app/protoc

# Install grpc
RUN go get google.golang.org/grpc

# PHP support
RUN apt-get -y install libz-dev

RUN git clone -b v1.34.0 https://github.com/grpc/grpc

WORKDIR /app/grpc
RUN git submodule update --init
RUN make grpc_php_plugin

# Install Go packages
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

# Seems to be renamed to protoc-gen-openapiv2
#RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# Added this for PHP server
RUN go get github.com/spiral/php-grpc/cmd/protoc-gen-php-grpc

ENV PATH="$PATH:/app/protoc/bin:/app/bin:/app"
RUN echo $PATH
RUN mkdir -p /app/gen/swagger
RUN mkdir -p /app/gen/php
RUN mkdir -p /app/gen/go
RUN mkdir -p /app/gen/swagger/search
WORKDIR /app
