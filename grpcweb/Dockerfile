FROM golang:1.12.1-alpine3.9

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/bin

RUN apk update && \
    apk upgrade && \
    apk add git

#install grpcurl
RUN go get -u github.com/fullstorydev/grpcurl && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl
    
RUN rm -rf /go/pkg

#install go pakage
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/grpc-ecosystem/go-grpc-middleware
RUN go get -u github.com/urfave/cli
RUN go get -u google.golang.org/grpc/credentials
RUN go get -u google.golang.org/grpc/reflection


