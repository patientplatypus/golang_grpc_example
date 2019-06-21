FROM golang:latest
# FROM grpc/go:latest

WORKDIR /
COPY . .

# install protobuf from source

RUN apt-get update
RUN apt-get -y install git unzip build-essential autoconf libtool
RUN git clone https://github.com/google/protobuf.git && \
    cd protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    ldconfig && \
    make clean && \
    cd .. && \
    rm -r protobuf

ENV GO111MODULE=on
RUN go mod init golang_backend

RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go
# RUN go get github.com/improbable-eng/grpc-web/go/grpcweb

# RUN go build main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
EXPOSE 8000
CMD ["./main"]


